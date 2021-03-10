package mailto

import (
	"bytes"
	"net/url"
	"strings"
	"text/template"

	"github.com/vivian-hua/civic-qa/services/mailto/internal/model"
)

// TODO: replies via https://tools.ietf.org/html/rfc6068#section-6.1

const (
	// InvalidMailto is returned if a mailto cannot be generated
	InvalidMailto = "INVALID MAILTO"

	// anchor mailto tag template
	mailtoTemplate = "<a href=\"mailto:{{.Addresses}}{{.Params}}\">{{.InnerText}}</a>"
)

// templateData is used by the template to construct
// the actual <a> tag
type templateData struct {
	Addresses string
	Params    string
	InnerText string
}

// Generate returns a mailto anchor tag for a given config
func Generate(request model.Request) ([]byte, error) {

	// Parse the template
	tmpl, err := template.New("mailto").Parse(mailtoTemplate)
	if err != nil {
		return []byte(InvalidMailto), err
	}

	params := buildParameters(request)

	// execute the template
	var buf bytes.Buffer
	err = tmpl.Execute(&buf,
		templateData{
			Addresses: strings.Join(request.To, ","),
			InnerText: request.InnerText,
			Params:    params,
		})

	if err != nil {
		return []byte(InvalidMailto), err
	}

	// return the resulting mailto anchor tag bytes
	return buf.Bytes(), nil
}

// buildParameters returns a string of mailto parameters
func buildParameters(request model.Request) string {
	var params []string
	if len(request.Cc) > 0 {
		params = append(params, "cc="+strings.Join(request.Cc, ","))
	}
	if len(request.Bcc) > 0 {
		params = append(params, "bcc="+strings.Join(request.Bcc, ","))
	}
	if request.Subject != "" {
		params = append(params, "subject="+url.PathEscape(request.Subject))
	}
	if request.Body != "" {
		params = append(params, "body="+url.PathEscape(request.Body))
	}

	// combine all parameters into one string
	paramStr := ""
	if len(params) > 0 {
		paramStr = "?" + strings.Join(params, "&")
	}

	return paramStr
}
