package mailto

import (
	"bytes"
	"net/url"
	"regexp"
	"strings"
	"text/template"
)

// TODO: replies via https://tools.ietf.org/html/rfc6068#section-6.1

const (
	// InvalidMailto is returned if a mailto cannot be generated
	InvalidMailto = "INVALID MAILTO"

	// anchor mailto tag template
	mailtoTemplate = "<a href=\"mailto:{{.Addresses}}{{.Params}}\">{{.InnerText}}</a>"
)

var (
	// email regex from: https://golangcode.com/validate-an-email-address/
	emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

// templateData is used by the template to construct
// the actual <a> tag
type templateData struct {
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

	params := buildParameters(config)

	// execute the template
	var buf bytes.Buffer
	err = tmpl.Execute(&buf,
		templateData{
			Addresses: strings.Join(config.To, ","),
			InnerText: config.InnerText,
			Params:    params,
		})

	if err != nil {
		return InvalidMailto, err
	}

	// return the resulting mailto anchor tag
	return buf.String(), nil
}

// buildParameters returns a string of mailto parameters
func buildParameters(config Config) string {
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

	return paramStr
}
