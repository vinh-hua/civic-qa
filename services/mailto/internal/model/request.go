package model

import (
	"errors"
	"fmt"
	"regexp"
)

var (
	defaultConfig = Request{To: []string{"mail@example.com"}}

	// email regex from: https://golangcode.com/validate-an-email-address/
	emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

// Request contains parameters to create a mailto tag
type Request struct {
	To        []string `json:"to"`
	Cc        []string `json:"cc"`
	Bcc       []string `json:"bcc"`
	Subject   string   `json:"subject"`
	Body      string   `json:"body"`
}

// Validate performs field validation on a Request, returning an error
// if validation fails
func (request Request) Validate() error {
	if len(request.To) < 1 {
		return errors.New("To must contain at least 1 address")
	}

	for i := range request.To {
		if !verifyEmail(request.To[i]) {
			return fmt.Errorf("Invalid Email Address in To: %s", request.To[i])
		}
	}

	for i := range request.Cc {
		if !verifyEmail(request.Cc[i]) {
			return fmt.Errorf("Invalid Email Address in Cc: %s", request.Cc[i])
		}
	}

	for i := range request.Bcc {
		if !verifyEmail(request.Bcc[i]) {
			return fmt.Errorf("Invalid Email Address in Bcc: %s", request.Bcc[i])
		}
	}

	return nil
}

// verifyEmail uses a regular expression
// to validate an email address
func verifyEmail(email string) bool {
	return emailRegex.MatchString(email)
}
