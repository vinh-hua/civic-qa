package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/team-ravl/civic-qa/service/account/internal/context"

	"github.com/team-ravl/civic-qa/services/common/config"
	aggregator "github.com/team-ravl/civic-qa/services/logAggregator/pkg/middleware"
)

const (
	// VersionBase is the API route base
	VersionBase = "/v0"
	// APIVersion is the API semantic version
	APIVersion = "v0.0.0"
)

func main() {
	// config
	var cfg config.Provider = &config.EnvProvider{}
	cfg.SetVerbose(true)

	// routers
	router := mux.NewRouter()
	api := router.PathPrefix(VersionBase).Subrouter()

	// middleware
	router.Use(aggregator.NewAggregatorMiddleware(&aggregator.Config{
		AggregatorAddress: cfg.GetOrFallback("AGG_ADDR", "http://localhost:8888"),
		ServiceName:       "account",
		SkipSuccesses:     true,
		StdoutErrors:      true,
		Timeout:           10 * time.Second,
	}))

	// handler context
	ctx, err := context.BuildContext(cfg)
	if err != nil {
		log.Fatalf("Could not create handler context: %v", err)
	}

	// routes
	api.Handle("/signup", http.HandlerFunc(ctx.HandleSignup))
	api.Handle("/login", http.HandlerFunc(ctx.HandleLogin))
	api.Handle("/logout", http.HandlerFunc(ctx.HandleLogout))
	api.Handle("/getsession", http.HandlerFunc(ctx.HandleGetSession))

	// start server
	addr := cfg.GetOrFallback("ADDR", ":8080")
	log.Printf("Server %s running on %s", APIVersion, addr)
	log.Fatal(http.ListenAndServe(addr, router))
}
