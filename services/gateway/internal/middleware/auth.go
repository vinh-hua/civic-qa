package middleware

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/vivian-hua/civic-qa/services/common/model"
)

const (
	getSessionPath      = "/getsession"
	authorizationHeader = "Authorization"
	xAuthUserIDHeader   = "X-AuthUser-ID"
)

// Config holds settings for an AuthMiddleware
type Config struct {
	AccountServiceURL string
}

type authError struct {
	message string
	status  int
}

// NewAuthMiddleware returns a gorilla mux.HandlerFunc middleware
// that authenticates all incoming requests and adds an X-AuthUser-ID
// header with the authenticated users ID.
// Responds with an error if authentication fails.
func NewAuthMiddleware(config *Config) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			// delete existing auth headers (sent by user)
			r.Header.Del(xAuthUserIDHeader)

			// authenticate the request
			state, authErr := authenticate(r, config)
			if authErr != nil {
				http.Error(w, authErr.message, authErr.status)
				return
			}

			// set auth header and continue
			r.Header.Set(xAuthUserIDHeader, fmt.Sprint(state.User.ID))
			h.ServeHTTP(w, r)
		})
	}
}

func authenticate(request *http.Request, config *Config) (*model.SessionState, *authError) {
	authReq, err := http.NewRequest(http.MethodGet, config.AccountServiceURL+getSessionPath, nil)
	if err != nil {
		log.Printf("Error creating getsession request: %v", err)
		return nil, &authError{message: "Internal Server Error", status: http.StatusInternalServerError}
	}

	// copy the headers
	for k := range request.Header {
		authReq.Header.Set(k, request.Header.Get(k))
	}

	// make the request to account
	client := http.Client{}
	resp, err := client.Do(authReq)
	if err != nil {
		log.Printf("Error making getsession request: %v", err)
		return nil, &authError{message: "Internal Server Error", status: http.StatusInternalServerError}
	}

	defer resp.Body.Close()

	// handle non-ok responses
	if resp.StatusCode >= 400 {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Error parsing getSession error: %v", err)
			return nil, &authError{message: "Internal Server Error", status: http.StatusInternalServerError}
		}
		return nil, &authError{message: string(bodyBytes), status: resp.StatusCode}
	}

	// decode the response
	var state model.SessionState
	err = json.NewDecoder(resp.Body).Decode(&state)
	if err != nil {
		log.Printf("Error parsing getsession response: %v", err)
		return nil, &authError{message: "Internal Server Error", status: http.StatusInternalServerError}
	}

	return &state, nil
}
