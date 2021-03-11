package context

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/team-ravl/civic-qa/services/common/parse"

	"github.com/gorilla/mux"
	"github.com/team-ravl/civic-qa/services/form/internal/model"
	"github.com/team-ravl/civic-qa/services/form/internal/repository/response"
)

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

	queryParams := r.URL.Query()

	query := response.Query{
		Name:         queryParams.Get("name"),
		EmailAddress: queryParams.Get("emailAddress"),
		InquiryType:  queryParams.Get("inquiryType"),
		Subject:      queryParams.Get("subject"),
		ActiveOnly:   parse.ParseBoolOrDefault(queryParams.Get("activeOnly")),
		TodayOnly:    parse.ParseBoolOrDefault(queryParams.Get("todayOnly")),
		FormID:       parse.ParseUintOrDefault(queryParams.Get("formID")),
	}

	// get the associated responses
	responses, err := ctx.ResponseStore.GetResponses(userID, query)
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

// HandlePatchResponse PATCH /responses/{responseID}
// Allows an authenticated user to update a responses 'open' state
func (ctx *Context) HandlePatchResponse(w http.ResponseWriter, r *http.Request) {

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

	// parse the request
	var patchResponse model.PatchResponse
	err = json.NewDecoder(r.Body).Decode(&patchResponse)
	if err != nil {
		http.Error(w, "Malformed Request", http.StatusBadRequest)
		return
	}

	// update the state
	err = ctx.ResponseStore.PatchByID(userID, uint(responseID), patchResponse.Active)
	if err != nil {
		log.Printf("Error updating FormResponse Open State: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// respond
	w.WriteHeader(http.StatusOK)
}
