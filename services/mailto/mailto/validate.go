package mailto

import (
	"errors"
	"fmt"
)

// validateConfig ensures To has at least 1 address
// and all addresses are valid
func validateConfig(config Config) error {
	if len(config.To) < 1 {
		return errors.New("To must contain at least 1 address")
	}

	for i := range config.To {
		if !verifyEmail(config.To[i]) {
			return fmt.Errorf("Invalid Email Address in To: %s", config.To[i])
		}
	}

	for i := range config.Cc {
		if !verifyEmail(config.Cc[i]) {
			return fmt.Errorf("Invalid Email Address in Cc: %s", config.Cc[i])
		}
	}

	for i := range config.Bcc {
		if !verifyEmail(config.Bcc[i]) {
			return fmt.Errorf("Invalid Email Address in Bcc: %s", config.Bcc[i])
		}
	}

	return nil

}

// verifyEmail uses a regular expression
// to validate an email address
func verifyEmail(email string) bool {
	return emailRegex.MatchString(email)
}
