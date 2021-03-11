package context

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/vivian-hua/civic-qa/services/form/internal/model"
	"github.com/vivian-hua/civic-qa/services/form/internal/repository/form"
)

// HandleCreateForm POST /forms
// Allows an authenticated user to create a new form
func (ctx *Context) HandleCreateForm(w http.ResponseWriter, r *http.Request) {
	// check content type
	if r.Header.Get("content-type") != "application/json" {
		http.Error(w, "Content type not allowed", http.StatusUnsupportedMediaType)
		return
	}

	// get the auth user
	userID, authErr := getAuthUserID(r)
	if authErr != nil {
		http.Error(w, authErr.message, authErr.code)
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
// Allows an authenticated user to view all their forms
func (ctx *Context) HandleGetForms(w http.ResponseWriter, r *http.Request) {

	// get the auth user
	userID, authErr := getAuthUserID(r)
	if authErr != nil {
		http.Error(w, authErr.message, authErr.code)
		return
	}

	// retrieve all their forms
	forms, err := ctx.FormStore.GetByUserID(userID)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// return forms
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&forms)
	if err != nil {
		log.Printf("Error encoding forms: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// HandleGetSpecificForm GET /forms/{formID}
// Allows an authenticated user to see a specific form
func (ctx *Context) HandleGetSpecificForm(w http.ResponseWriter, r *http.Request) {
	// parse the URL parameter
	vars := mux.Vars(r)
	formID, err := strconv.ParseUint(vars["formID"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid Form ID", http.StatusBadRequest)
		return
	}

	// get the auth user
	userID, authErr := getAuthUserID(r)
	if authErr != nil {
		http.Error(w, authErr.message, authErr.code)
		return
	}

	// retrieve the form
	formData, err := ctx.FormStore.GetByID(uint(formID))
	if err != nil {
		if err == form.ErrFormNotFound {
			http.Error(w, "Form Not Found", http.StatusNotFound)
			return
		}
		log.Printf("Error retrieving form: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// make sure this form belongs to the user
	if formData.UserID != userID {
		http.Error(w, "This form does not belong to you", http.StatusForbidden)
		return
	}

	// return the form
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&formData)
	if err != nil {
		log.Printf("Error encoding forms: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
