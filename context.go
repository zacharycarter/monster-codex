package main

type Context interface {
	Name() string
}

type TemplateContext struct {
	TemplateName string
	Errors       []string
	Notices      []string
}

func (ctx *TemplateContext) Name() string {
	return ctx.TemplateName
}

type HomeContext struct {
	*TemplateContext
	Monsters []Monster
}
