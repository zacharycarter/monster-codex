package template

import (
	"html/template"
	"net/http"

	"github.com/zacharycarter/monster-codex/util"
)

type Context struct {
}

var templateCache = template.Must(template.ParseGlob("./template/*.html"))

func RenderTemplate(w http.ResponseWriter, templates []string) {
	for i := range templates {
		err := templateCache.ExecuteTemplate(w, templates[i], Context{})
		util.CheckErr(err, "Error rendering template: ")
	}
}
