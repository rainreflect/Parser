package handler

import (
	"html/template"
	"net/http"

	"github.com/rainreflect/parser/domain"
)

func ArticlesHandler(articles []domain.Items, tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := tmpl.Execute(w, articles); err != nil {
			http.Error(w, "Unable to render template", http.StatusInternalServerError)
		}
	}
}
