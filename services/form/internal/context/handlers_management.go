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
	"github.com/vivian-hua/civic-qa/services/form/internal/repository/response"
)

const (
	// xAuthUserIDHeader is the header where authenticated userID should be
	xAuthUserIDHeader = "X-AuthUser-ID"
)

// authError contains information about
// a missing or invalid xAuthUserIDHeader
type authError struct {
	message string
	code    int
}

// getAuthUserID returns the userID of a requests authenticated user if present,
// or returns an authError
func getAuthUserID(r *http.Request) (uint, *authError) {
	// check that the user is authenticated
	authUserStr := r.Header.Get(xAuthUserIDHeader)
	if authUserStr == "" {
		return 0, &authError{message: "No Authorization Found", code: http.StatusUnauthorized}
	}

	// parse userID
	userID, err := strconv.ParseUint(authUserStr, 10, 64)
	if err != nil {
		log.Printf("Error parsing userID: %v", err)
		return 0, &authError{message: "Invalid Authorization", code: http.StatusUnauthorized}
	}

	return uint(userID), nil
}

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

// HandleGetFormResponses GET /form/{formID}/responses
// Allows an authenticated user to see response to a specific form
func (ctx *Context) HandleGetFormResponses(w http.ResponseWriter, r *http.Request) {
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

	// retrieve the responses
	responses, err := ctx.ResponseStore.GetByFormID(uint(formID))
	if err != nil {
		log.Printf("Error retrieving responses: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// return the responses
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&responses)
	if err != nil {
		log.Printf("Error encoding forms: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// HandleGetSpecificResponse GET /responses/{responseID}
// Allows an authenticated user to view a specific response
func (ctx *Context) HandleGetSpecificResponse(w http.ResponseWriter, r *http.Request) {
	// parse the URL parameter
	vars := mux.Vars(r)
	responseID, err := strconv.ParseUint(vars["responseID"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid response ID", http.StatusBadRequest)
		return
	}

	// get the auth user
	userID, authErr := getAuthUserID(r)
	if authErr != nil {
		http.Error(w, authErr.message, authErr.code)
		return
	}

	// retrieve the response
	respData, err := ctx.ResponseStore.GetByID(uint(responseID))
	if err != nil {
		if err == response.ErrResponseNotFound {
			http.Error(w, "Response not found", http.StatusNotFound)
			return
		}
		log.Printf("Error retrieving response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// get the associated form
	formData, err := ctx.FormStore.GetByID(respData.FormID)
	if err != nil {
		log.Printf("Error retrieving form: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// ensure the form belongs to the user
	if formData.UserID != userID {
		http.Error(w, "This response does not belong to you", http.StatusForbidden)
		return
	}

	// return the response
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&respData)
	if err != nil {
		log.Printf("Error encoding forms: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// HandleGetResponses GET /responses
// Allows an authenticated user to get responses to all their forms
func (ctx *Context) HandleGetResponses(w http.ResponseWriter, r *http.Request) {
	// get the auth user
	userID, authErr := getAuthUserID(r)
	if authErr != nil {
		http.Error(w, authErr.message, authErr.code)
		return
	}

	// get the associated responses
	responses, err := ctx.ResponseStore.GetByUserID(userID)
	if err != nil {
		log.Printf("Error retrieving responses by userID: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// return the responses
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&responses)
	if err != nil {
		log.Printf("Error encoding forms: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
