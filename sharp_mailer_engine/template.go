package sharp_mailer_engine

import (
	"bytes"
	"fmt"
	"html/template"
	"path/filepath"
)

func (e *SharpMailerEngine) LoadTemplates(dir string) error {
	t, err := template.ParseGlob(filepath.Join(dir, "*.html"))
	if err != nil {
		return err
	}

	e.templateCache = t
	return nil
}

func (e *SharpMailerEngine) parse(key TemplateKey, data interface{}) (string, error) {
	var buf bytes.Buffer
	err := e.templateCache.ExecuteTemplate(&buf, string(key), data)
	if err != nil {
		return "", fmt.Errorf("mailer: execution error for %s: %w", key, err)
	}
	return buf.String(), nil
}
