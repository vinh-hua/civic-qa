package context

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/vivian-hua/civic-qa/services/form/internal/model"
)

const (
	xAuthUserIDHeader = "X-AuthUser-ID"
)

type authError struct {
	Message string
	Code    int
}

func getAuthUserID(r *http.Request) (uint, *authError) {
	// check that the user is authenticated
	authUserStr := r.Header.Get(xAuthUserIDHeader)
	if authUserStr == "" {
		return 0, &authError{Message: "No Authorization Found", Code: http.StatusUnauthorized}
	}

	// parse userID
	userID, err := strconv.ParseUint(authUserStr, 10, 64)
	if err != nil {
		log.Printf("Error parsing userID: %v", err)
		return 0, &authError{Message: "Internal Server Error", Code: http.StatusInternalServerError}
	}

	return uint(userID), nil
}

// HandleCreateForm POST /forms
func (ctx *Context) HandleCreateForm(w http.ResponseWriter, r *http.Request) {
	// check content type
	if r.Header.Get("content-type") != "application/json" {
		http.Error(w, "Content type not allowed", http.StatusUnsupportedMediaType)
		return
	}

	// get the auth user
	userID, authErr := getAuthUserID(r)
	if authErr != nil {
		http.Error(w, authErr.Message, authErr.Code)
		return
	}

	// parse request
	var newForm model.NewFormRequest
	err := json.NewDecoder(r.Body).Decode(&newForm)
	if err != nil {
		http.Error(w, "Malformed Request", http.StatusBadRequest)
		return
	}

	// create the form
	form := &model.Form{
		Name:      newForm.Name,
		CreatedAt: time.Now(),
		UserID:    uint(userID),
	}

	err = ctx.FormStore.Create(form)
	if err != nil {
		log.Printf("Error creating form: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// respond with the created form
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(&form)
	if err != nil {
		log.Printf("Error encoding form: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// HandleGetForms GET /forms
func (ctx *Context) HandleGetForms(w http.ResponseWriter, r *http.Request) {

	// get the auth user
	userID, authErr := getAuthUserID(r)
	if authErr != nil {
		http.Error(w, authErr.Message, authErr.Code)
		return
	}

	// retrieve all their forms
	forms, err := ctx.FormStore.GetByUserID(userID)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// return forms
	// respond with the created form
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&forms)
	if err != nil {
		log.Printf("Error encoding forms: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

}
