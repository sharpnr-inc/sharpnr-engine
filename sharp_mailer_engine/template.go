package sharp_mailer_engine

import (
	"bytes"
	"fmt"
	"html/template"
	"path/filepath"
)

// LoadTemplates parses all *.html files in dir into the template cache.
// Must be called before Start().
func (e *SharpMailerEngine) LoadTemplates(dir string) error {
	t, err := template.ParseGlob(filepath.Join(dir, "*.html"))
	if err != nil {
		return fmt.Errorf("mailer: failed to load templates from %s: %w", dir, err)
	}
	e.templateCache = t
	return nil
}

// parse renders a named template with data and returns the HTML string.
func (e *SharpMailerEngine) parse(key TemplateKey, data any) (string, error) {
	if e.templateCache == nil {
		return "", fmt.Errorf("mailer: template cache is nil — did you call LoadTemplates?")
	}

	var buf bytes.Buffer
	// Go template name = the {{define "name"}} string inside the file,
	// which for base_layout.html is "base_layout.html" itself
	if err := e.templateCache.ExecuteTemplate(&buf, string(key), data); err != nil {
		return "", fmt.Errorf("mailer: execute %s: %w", key, err)
	}

	return buf.String(), nil
}
