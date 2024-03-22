package web

import (
	"book_ex/internal/models"
	"html/template"
	"path/filepath"
)

type templateData struct {
	Item  *models.Item
	Items []*models.Item
}

func NewTemplateCache() (map[string]*template.Template, error) {

	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./ui/templates/pages/*.gohtml")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		t, err := template.ParseFiles("./ui/templates/base.gohtml")
		if err != nil {
			return nil, err
		}

		t, err = t.ParseGlob("./ui/templates/partials/*.gohtml")
		if err != nil {
			return nil, err
		}

		t, err = t.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		cache[name] = t
	}

	return cache, nil

}
