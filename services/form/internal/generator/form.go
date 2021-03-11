package generator

import (
	"bytes"
	"html/template"
	"log"

	"github.com/team-ravl/civic-qa/services/form/internal/model"
)

var (
	// formTemplate is an HTML template to a form.
	// Fatal if does not exist or not parsable
	formTemplate = mustParse("/templates/form.html")
)

// mustParse attempts to parse templates in given files.
// returns the parsed template or fatal
func mustParse(filenames ...string) *template.Template {
	formTemplate, err := template.ParseFiles(filenames...)
	if err != nil {
		log.Fatalf("Failed to parse template file %s, %v", filenames, err)
		return nil
	}

	return formTemplate
}

// GetHTML returns bytes for the HTML form of a given Form
func GetHTML(form *model.Form) ([]byte, error) {

	var buffer bytes.Buffer
	err := formTemplate.Execute(&buffer, form)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
