package main

import (
	"html/template"
	"io/fs"
	"path/filepath"
	"time"

	"snippetbox.sagyzdop.com/internal/models"
	"snippetbox.sagyzdop.com/ui"
)

type templateData struct {
	CurrentYear     int
	Snippet         models.Snippet
	Snippets        []models.Snippet
	Form            any
	Flash           string
	IsAuthenticated bool
	CSRFToken       string
}

var functions = template.FuncMap{
	"humanDate": humanDate,
}

func humanDate(t time.Time) string {
	// Return the empty string if time has the zero value.
	if t.IsZero() {
		return ""
	}
	// Convert the time to UTC before formatting it.
	return t.UTC().Format("02 Jan 2006 at 15:04")
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := fs.Glob(ui.Files, "html/pages/*.tmpl.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		patterns := []string{
			"html/base.tmpl.html",
			"html/partials/*.tmpl.html",
			page,
		}
		// Parse the base template file into a template set.
		ts, err := template.New(name).Funcs(functions).ParseFS(ui.Files, patterns...)
		if err != nil {
			return nil, err
		}

		// Add the template set to the map as normal...
		cache[name] = ts
	}

	return cache, nil
}
