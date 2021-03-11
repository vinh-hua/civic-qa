package context

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/team-ravl/civic-qa/service/account/internal/model"
)

const (
	minPasswordLength = 8
	maxPasswordLength = 55
)

var (
	// email regex from: https://golangcode.com/validate-an-email-address/
	emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	errInvalidEmail = errors.New("Invalid Email Address")

	errPasswordNoMatch = errors.New("Passwords must match")
	errPassTooShort    = fmt.Errorf("Password too short, must be at least %d characters", minPasswordLength)
	errPassTooLong     = fmt.Errorf("Password too long, must not exceed %d characters", maxPasswordLength)
)

// validates a NewUser
// ensures email is valid, password and confirmation matches,
// and password length is within bounds
func validateNewUser(newUser model.NewUserRequest) error {
	// validate email
	if !emailRegex.MatchString(newUser.Email) {
		return errInvalidEmail
	}

	// check if passwords match
	if newUser.Password != newUser.PasswordConfirm {
		return errPasswordNoMatch
	}

	// ensure len(password) >= minimum
	if len(newUser.Password) < minPasswordLength {
		return errPassTooShort
	}

	// ensure len(password) <= maximum
	if len(newUser.Password) > maxPasswordLength {
		return errPassTooLong
	}

	return nil
}
