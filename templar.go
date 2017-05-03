package templar

import (
	"bytes"
	"io/ioutil"
	"text/template"
)

// CreateFileByTemplate parses template, substitutes vars and saves file
func CreateFileByTemplate(templateFile, destFile string, params interface{}) error {
	template, err := ParseTemplateFile(templateFile, params)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(destFile, []byte(template), 0644)
}

// ParseTemplateFile parses a template file and supstitutes the variables, returns template instance with variables replaced
func ParseTemplateFile(templateFile string, params interface{}) (string, error) {
	tplFile, err := ioutil.ReadFile(templateFile)

	if err != nil {
		return "", err
	}

	r, err := ParseTemplateString(string(tplFile), params)
	return r, err
}

// ParseTemplateString parses a template string and supstitutes the variables, returns template instance with variables replaced
func ParseTemplateString(templateString string, params interface{}) (string, error) {
	t := template.Must(template.New("letter").Parse(templateString))

	var doc bytes.Buffer
	errParse := t.Execute(&doc, params)
	if errParse != nil {
		return "", errParse
	}
	resp := doc.String()

	return resp, nil
}
