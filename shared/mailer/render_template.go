package mailer

import (
	"bytes"
	"html/template"
)

func renderTemplate(path string, data any) (string, error) {
	tpl, err := template.ParseFiles(path)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := tpl.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}
