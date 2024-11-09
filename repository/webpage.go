package repository

import (
	"html/template"
	"log"
	"net/http"
)

type WebPageData struct {
	Title  string
	Script template.JS
}

type WebPage struct {
	Data     *WebPageData
	Template *template.Template
}

func InitWebPageRepo(data *WebPageData, template *template.Template) *WebPage {
	return &WebPage{Data: data, Template: template}
}

func (repo *WebPage) Render(w http.ResponseWriter, templateName string, title string) {
	repo.Data.Title = title
	if err := repo.Template.ExecuteTemplate(w, templateName, repo.Data); err != nil {
		log.Fatalf("Error executing template: %v", err)
	}
}
