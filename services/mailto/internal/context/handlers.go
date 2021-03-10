package context

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/vivian-hua/civic-qa/services/mailto/internal/mailto"
	"github.com/vivian-hua/civic-qa/services/mailto/internal/model"
)

// Mailto POST /mailto
// Generates a mailto anchor tag given a valid request
func (ctx *Context) Mailto(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("content-type") != "application/json" {
		http.Error(w, "Expected content-type: application/json", http.StatusUnsupportedMediaType)
		return
	}

	// parse request
	var request model.Request
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Malformed Request", http.StatusBadRequest)
		return
	}

	// validate the request
	err = request.Validate()
	if err != nil {
		http.Error(w, fmt.Sprintf("Malformed Request: %s", err.Error()), http.StatusBadRequest)
		return
	}

	// generate the HTML
	htmlBytes, err := mailto.Generate(request)
	if err != nil {
		log.Printf("Error generating HTML: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// return the HTML
	w.Header().Set("content-type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(htmlBytes)
	if err != nil {
		log.Printf("Error writing: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
