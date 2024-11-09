package router

import (
	"database/sql"
	"github.com/go-chi/chi/v5"
	"html/template"
	"log"
	"net/http"
	"project/handler"
	"project/repository"
	"project/service"
)

func initTemplate() (*repository.WebPageData, *template.Template) {
	tmpl, err := template.ParseGlob("view/*.html")
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
		return nil, nil
	}

	return &repository.WebPageData{}, tmpl
}

func NewRouter(db *sql.DB) *chi.Mux {
	r := chi.NewRouter()

	handleCategory := handler.InitCategoryHandler(*service.InitCategoryService(*repository.InitCategoryRepo(db)))
	handleItem := handler.InitItemHandler(*service.InitItemService(*repository.InitItemRepo(db)))
	handleWebTemplate := handler.InitWebPageHandler(*service.InitWebPageService(*repository.InitWebPageRepo(initTemplate())))

	r.Route("/api", func(r chi.Router) {
		r.Route("/categories", func(r chi.Router) {
			r.Get("/", func(w http.ResponseWriter, r *http.Request) {})
			r.Post("/", handleCategory.Create)
		})

		r.Route("/items", func(r chi.Router) {
			r.Get("/", func(w http.ResponseWriter, r *http.Request) {})
			r.Post("/", handleItem.Create)
		})

	})

	r.Get("/app.css", handleWebTemplate.Static)

	return r
}
