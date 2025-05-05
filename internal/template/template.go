package template

import (
	"html/template"
	"io/fs"
	"log/slog"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/obzva/dngyng1000/ui"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Template struct {
	Cache map[string]*template.Template
}

func New() (*Template, error) {
	t := &Template{
		Cache: make(map[string]*template.Template),
	}

	pages, err := fs.Glob(ui.Files, "tmpls/pages/*.tmpl")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).ParseFS(ui.Files, "tmpls/base.tmpl", page)
		if err != nil {
			return nil, err
		}

		t.Cache[name] = ts
	}

	return t, nil
}

func (t *Template) Render(logger *slog.Logger, w http.ResponseWriter, r *http.Request, name string) {
	ts, ok := t.Cache[name]
	if !ok {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	caser := cases.Title(language.English)
	pageTitle := caser.String(strings.TrimSuffix(name, ".tmpl"))

	err := ts.ExecuteTemplate(w, "base", struct {
		PageTitle string
	}{
		PageTitle: pageTitle,
	})
	if err != nil {
		logger.Error(err.Error())
	}
}
