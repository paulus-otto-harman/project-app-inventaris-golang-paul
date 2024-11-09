package service

import (
	"net/http"
	"project/repository"
)

type WebPageService struct {
	WebPage repository.WebPage
}

func InitWebPageService(repo repository.WebPage) *WebPageService {
	return &WebPageService{WebPage: repo}
}

func (repo *WebPageService) Render(w http.ResponseWriter, templateName string, title string) {
	repo.WebPage.Render(w, templateName, title)
}
