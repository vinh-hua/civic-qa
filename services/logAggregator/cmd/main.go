package main

import (
	// standard
	"log"
	"net/http"

	// 3rd party
	"github.com/gorilla/mux"

	// common
	"github.com/vivian-hua/civic-qa/services/common/environment"

	// internal
	"github.com/vivian-hua/civic-qa/services/logAggregator/internal/handlercontext"
	"github.com/vivian-hua/civic-qa/services/logAggregator/internal/models"
)

const (
	// VersionBase is the API route base
	VersionBase = "/v0"
	// APIVersion is the API semantic version
	APIVersion = "v0.0.0"
)

var (
	// Environment
	addr   = environment.GetEnvOrFallback("ADDR", ":8888")
	dbPath = environment.GetEnvOrFallback("DBPATH", "logs.db")
)

func main() {

	// Routers
	router := mux.NewRouter()
	api := router.PathPrefix(VersionBase).Subrouter()

	// Handler context
	store := models.NewSQLiteLogStore("logs.db", true)
	ctx := &handlercontext.Context{Store: store}

	// Routes
	api.HandleFunc("/log", ctx.HandleLog)
	api.HandleFunc("/query", ctx.HandleQuery)

	// Start Server
	log.Printf("Server %s running on %s", APIVersion, addr)
	log.Fatal(http.ListenAndServe(addr, router))
}
