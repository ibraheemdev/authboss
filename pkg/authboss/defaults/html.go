package defaults

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"path"
	"path/filepath"

	"github.com/ibraheemdev/authboss/pkg/authboss"
)

// HTMLRenderer is a simple template renderer that stores a map of templates
type HTMLRenderer struct {
	// url mount path
	mountPath string

	// path to templates folder
	templatesPath string

	templates map[string]*template.Template
}

// NewHTMLRenderer :
func NewHTMLRenderer(mountPath, templatesPath string) *HTMLRenderer {
	return &HTMLRenderer{
		mountPath:     mountPath,
		templates:     make(map[string]*template.Template),
		templatesPath: templatesPath,
	}
}

// Render a page with data
func (h *HTMLRenderer) Render(ctx context.Context, page string, data authboss.HTMLData) (output []byte, contentType string, err error) {
	tmpl, ok := h.templates[page]
	if !ok {
		return nil, "", fmt.Errorf("the template %s does not exist", page)
	}
	buf := &bytes.Buffer{}
	err = tmpl.ExecuteTemplate(buf, filepath.Base(page), data)
	if err != nil {
		return nil, "", fmt.Errorf("failed to render template for page %s: %w", page, err)
	}
	return buf.Bytes(), "text/html", nil
}

// Load a template directory
func (h *HTMLRenderer) Load(templates ...string) error {
	funcMap := template.FuncMap{
		"mountpathed": func(location string) string {
			return path.Join(h.mountPath, location)
		},
	}

	for _, tpl := range templates {
		filePath := fmt.Sprintf("%s%s", h.templatesPath, tpl)
		template, err := template.New(tpl).Funcs(funcMap).ParseFiles(filePath)
		if err != nil {
			return fmt.Errorf("Could not parse template: %s", tpl)
		}
		h.templates[tpl] = template
	}
	return nil
}
