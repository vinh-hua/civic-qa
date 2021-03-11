package context

import (
	"log"
	"net/http"
	"strconv"
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
