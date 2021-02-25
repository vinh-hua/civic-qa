package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/vivian-hua/civic-qa/service/account/internal/context"
	"github.com/vivian-hua/civic-qa/services/common/environment"

	aggregator "github.com/vivian-hua/civic-qa/services/logAggregator/pkg/middleware"
)

const (
	// VersionBase is the API route base
	VersionBase = "/v0"
	// APIVersion is the API semantic version
	APIVersion = "v0.0.0"
)

var (
	addr           = environment.GetEnvOrFallback("ADDR", ":8080")
	aggregatorAddr = environment.GetEnvOrFallback("ADDR", ":8888")
)

func main() {
	// routers
	router := mux.NewRouter()
	api := router.PathPrefix(VersionBase).Subrouter()

	// middleware
	router.Use(aggregator.NewAggregatorMiddleware(&aggregator.Config{
		AggregatorAddress: aggregatorAddr,
		ServiceName:       "account",
		StdoutErrors:      true,
		Timeout:           10 * time.Second,
	}))

	// handler context
	ctx, err := context.BuildContext()
	if err != nil {
		log.Fatalf("Could not create handler context: %v", err)
	}

	// routes
	api.Handle("/signup", http.HandlerFunc(ctx.HandleSignup))
	api.Handle("/login", http.HandlerFunc(ctx.HandleLogin))
	api.Handle("/logout", http.HandlerFunc(ctx.HandleLogout))
	api.Handle("/getsession", http.HandlerFunc(ctx.HandleGetSession))

	// start server
	log.Printf("Server %s running on %s", APIVersion, addr)
	log.Fatal(http.ListenAndServe(addr, router))
}
