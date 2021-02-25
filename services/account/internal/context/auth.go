package context

import (
	"net/http"

	"github.com/vivian-hua/civic-qa/service/account/internal/model"
	"github.com/vivian-hua/civic-qa/service/account/internal/repository/session"
	common "github.com/vivian-hua/civic-qa/services/common/model"
	"golang.org/x/crypto/bcrypt"
)

const (
	// bcrypt hash cost factor
	bcryptCost = 10

	// Authorization header details
	authorizationHeader = "Authorization"
	authorizationSchema = "Bearer "
)

// hashPassword hashes a given password and
// returns the resulting byte-slice
func hashPassword(password string) ([]byte, error) {
	result, err := bcrypt.GenerateFromPassword([]byte(password), bcryptCost)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// checkPassword returns an error if Users password does not match LoginRequest
func checkPassword(user common.User, login model.LoginRequest) error {
	return bcrypt.CompareHashAndPassword(user.PassHash, []byte(login.Password))
}

// addAuthorizationHeader adds an authorization token
// to a responses headers for a given session.Token
func addAuthorizationHeader(w http.ResponseWriter, token session.Token) {
	w.Header().Set(authorizationHeader, authorizationSchema+string(token))
}
