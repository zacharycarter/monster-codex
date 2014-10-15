package main

import (
	"html/template"
	"io"
)

type TemplateRenderer interface {
	RenderTemplate(wr io.Writer, ctx Context) error
}

type reloadingTemplateRenderer struct{}

var templates *template.Template

func NewTemplateRenderer() *reloadingTemplateRenderer {
	return &reloadingTemplateRenderer{}
}

func (r reloadingTemplateRenderer) RenderTemplate(wr io.Writer, ctx Context) error {
	templates, err := template.ParseGlob("./templates/*.html")

	if err != nil {
		return err
	}

	return renderWithLayout(templates, wr, ctx.Name(), ctx)
}

func renderWithLayout(t *template.Template, wr io.Writer, name string, data interface{}) error {
	for _, templateName := range []string{"header.html", name, "footer.html"} {
		err := t.ExecuteTemplate(wr, templateName, data)

		if err != nil {
			return err
		}
	}

	return nil
}
