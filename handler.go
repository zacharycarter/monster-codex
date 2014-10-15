package main

import (
	"log"
	"net/http"
)

const (
	serverError string = "Internal server error"
)

type Handler interface {
	Handle(res http.ResponseWriter, req *http.Request)
}

type HomeHandler struct {
	TemplateRenderer
	Monsters Repository
}

func NewHomeHandler(renderer TemplateRenderer, repository Repository) *HomeHandler {
	return &HomeHandler{renderer, repository}
}

func (h HomeHandler) Context(res http.ResponseWriter, req *http.Request, templateName string) *TemplateContext {
	return &TemplateContext{TemplateName: templateName}
}

func (h HomeHandler) RenderPage(res http.ResponseWriter, req *http.Request, pageContext Context) {
	err := h.RenderTemplate(res, pageContext)

	if err != nil {
		log.Println("Error when rendering template:", err, ". Context:", pageContext)
	}
}

func (h *HomeHandler) Handle(res http.ResponseWriter, req *http.Request) {
	ctx := HomeContext{h.Context(res, req, "home.html"), nil}
	monsters, err := h.Monsters.All()

	if err != nil {
		log.Println("Error loading monsters from repository:", err)
		//ctx.AddError(errors.New("Unable to load monsters"))
	} else {
		ctx.Monsters = monsters
	}

	h.RenderPage(res, req, &ctx)
}
