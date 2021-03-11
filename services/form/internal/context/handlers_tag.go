package context

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/vivian-hua/civic-qa/services/form/internal/model"
)

/*
	/responses/{responseID}/tags:
		GET:
			get a list of tags on a given response

		DELETE:
			remove a tag to a response
			schema:
				{
					value: string
				}
*/

// HandleGetAllTags GET /tags
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

	// get all the tag values
	tagList := make([]string, len(tags))
	for i := range tags {
		tagList[i] = tags[i].Value
	}

	// marshal tag values to json
	tagsJSON, err := json.Marshal(&tagList)
	if err != nil {
		log.Printf("Error marshalling tags: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// write the response
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(tagsJSON)
	if err != nil {
		log.Printf("Error writing tags: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// HandleGetTags GET /responses/{responseID}/tags
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

	// get all the tag values
	tagList := make([]string, len(tags))
	for i := range tags {
		tagList[i] = tags[i].Value
	}

	// marshal tag values to json
	tagsJSON, err := json.Marshal(&tagList)
	if err != nil {
		log.Printf("Error marshalling tags: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// write the response
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(tagsJSON)
	if err != nil {
		log.Printf("Error writing tags: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// HandlePostTag POST /responses/{responseID}/tags
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
