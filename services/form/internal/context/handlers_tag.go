package context

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/team-ravl/civic-qa/services/form/internal/model"
)

// HandleGetAllTags GET /tags
// Allows an authenticated user to get a list of all tags on all responses
func (ctx *Context) HandleGetAllTags(w http.ResponseWriter, r *http.Request) {
	// get the auth user
	userID, authErr := getAuthUserID(r)
	if authErr != nil {
		http.Error(w, authErr.message, authErr.code)
		return
	}

	// get the tags
	tags, err := ctx.TagStore.GetAll(userID)
	if err != nil {
		log.Printf("Error getting all tags: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// write the response
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&tags)
	if err != nil {
		log.Printf("Error writing tags: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// HandleGetTags GET /responses/{responseID}/tags
// Allows an authenticated user to get all tags to a specific response
func (ctx *Context) HandleGetTags(w http.ResponseWriter, r *http.Request) {
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

	// Get the tags for the response
	tags, err := ctx.TagStore.GetByResponseID(userID, uint(responseID))
	if err != nil {
		log.Printf("Error getting tags by id: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// write the response
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&tags)
	if err != nil {
		log.Printf("Error writing tags: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// HandlePostTag POST /responses/{responseID}/tags
// Allows an authenticated user to add a tag to a respoonse
func (ctx *Context) HandlePostTag(w http.ResponseWriter, r *http.Request) {
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
	var request model.TagRequest
	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Malformed Request", http.StatusBadRequest)
		return
	}

	// create the tag
	err = ctx.TagStore.Create(userID, uint(responseID), request.Value)
	if err != nil {
		log.Printf("Error creating tag: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// write the response
	w.WriteHeader(http.StatusCreated)
}

// HandleDeleteTag DELETE /responses/{responseID}/tags
// Allows an authenticated user to remove a tag from a response
func (ctx *Context) HandleDeleteTag(w http.ResponseWriter, r *http.Request) {
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
	var request model.TagRequest
	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Malformed Request", http.StatusBadRequest)
		return
	}

	// delete the tag
	err = ctx.TagStore.Delete(userID, uint(responseID), request.Value)
	if err != nil {
		log.Printf("Error deleting tag: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// write the response
	w.WriteHeader(http.StatusOK)
}
