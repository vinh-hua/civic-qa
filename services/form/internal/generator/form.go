package generator

import (
	"bytes"
	"html/template"
	"log"

	"github.com/vivian-hua/civic-qa/services/form/internal/model"
)

var (
	formTemplate = mustParse("/templates/form.html")
)

func mustParse(filenames ...string) *template.Template {
	formTemplate, err := template.ParseFiles(filenames...)
	if err != nil {
		log.Fatalf("Failed to parse template file %s, %v", filenames, err)
	}

	return formTemplate
}

func GetHTML(form *model.Form) ([]byte, error) {

	var buffer bytes.Buffer
	err := formTemplate.Execute(&buffer, form)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
