package mailto

import (
	"bytes"
	"errors"
	"fmt"
	"net/url"
	"regexp"
	"strings"
	"text/template"
)

// TODO: replies via https://tools.ietf.org/html/rfc6068#section-6.1

const (
	// InvalidMailto is returned if a mailto cannot be generated
	InvalidMailto = "INVALID MAILTO"

	mailtoTemplate = "<a href=\"mailto:{{.Addresses}}{{.Params}}\">{{.InnerText}}</a>"
)

var (
	defaultConfig = Config{To: []string{"mail@example.com"}, InnerText: "Click Me!"}

	// email regex from: https://golangcode.com/validate-an-email-address/
	emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

// Config contains parameters to create a mailto tag
type Config struct {
	To        []string
	Cc        []string
	Bcc       []string
	Subject   string
	Body      string
	InnerText string
}

// executor is used by the template to construct
// the actual <a> tag
type executor struct {
	Addresses string
	Params    string
	InnerText string
}

// Generate returns a mailto anchor tag for a given config
func Generate(config Config) (string, error) {

	// Validate addresses
	err := validateConfig(config)
	if err != nil {
		return InvalidMailto, err
	}

	// Parse the template
	tmpl, err := template.New("mailto").Parse(mailtoTemplate)
	if err != nil {
		return InvalidMailto, err
	}

	// Create a list of optional parameters
	var params []string
	if len(config.Cc) > 0 {
		params = append(params, "cc="+strings.Join(config.Cc, ","))
	}
	if len(config.Bcc) > 0 {
		params = append(params, "bcc="+strings.Join(config.Bcc, ","))
	}
	if config.Subject != "" {
		params = append(params, "subject="+url.PathEscape(config.Subject))
	}
	if config.Body != "" {
		params = append(params, "body="+url.PathEscape(config.Body))
	}

	// combine all parameters into one string
	paramStr := ""
	if len(params) > 0 {
		paramStr = "?" + strings.Join(params, "&")
	}

	// execute the template
	var buf bytes.Buffer
	err = tmpl.Execute(&buf,
		executor{
			Addresses: strings.Join(config.To, ","),
			InnerText: config.InnerText,
			Params:    paramStr,
		})

	if err != nil {
		return InvalidMailto, err
	}

	// return the resulting mailto anchor tag
	return buf.String(), nil
}

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
