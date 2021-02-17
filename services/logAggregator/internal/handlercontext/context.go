package handlercontext

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/vivian-hua/civic-qa/services/logAggregator/internal/models"
)

// Context stores handler context information for logAggregator
type Context struct {
	Store models.LogStore
}

// HandleLog handles a Log request
func (ctx *Context) HandleLog(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Expected: POST", http.StatusMethodNotAllowed)
		return
	}

	// Parse request
	var logReq models.LogEntry
	err := json.NewDecoder(r.Body).Decode(&logReq)
	if err != nil {
		log.Printf("Failed to parse Log request: %v\n", err)
		http.Error(w, "Malformed Request", http.StatusBadRequest)
		return
	}

	// Create log entry
	logErr := ctx.Store.Log(logReq)
	if err != nil {
		log.Printf("Failed to Log: %v\n", logErr.Err)
		http.Error(w, "Failed to log", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// HandleQuery handles a query request
func (ctx *Context) HandleQuery(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Expected: POST", http.StatusMethodNotAllowed)
		return
	}

	// Parse request
	var queryReq models.LogQuery
	err := json.NewDecoder(r.Body).Decode(&queryReq)
	if err != nil {
		log.Printf("Failed to parse query request: %v\n", err)
		http.Error(w, "Malformed Request", http.StatusBadRequest)
		return
	}

	// Query
	result, queryErr := ctx.Store.Query(queryReq)
	if err != nil {
		log.Printf("Failed to query: %v\n", queryErr.Err)
		http.Error(w, "Failed to query", http.StatusBadRequest)
		return
	}

	// Return results
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&result)
}
