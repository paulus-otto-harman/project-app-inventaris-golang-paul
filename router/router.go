package router

import (
	"database/sql"
	"github.com/go-chi/chi/v5"
	"html/template"
	"log"
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
	handleInvestment := handler.InitInvestmentHandler(*service.InitInvestmentService(*repository.InitInvestmentRepo(db)))
	handleWebTemplate := handler.InitWebPageHandler(*service.InitWebPageService(*repository.InitWebPageRepo(initTemplate())))

	r.Route("/api", func(r chi.Router) {
		r.Route("/categories", func(r chi.Router) {
			r.Post("/", handleCategory.Create)
			r.Get("/", handleCategory.All)
			r.Get("/{id}", handleCategory.Get)
			r.Put("/{id}", handleCategory.Update)
			r.Delete("/{id}", handleCategory.Delete)
		})

		r.Route("/items", func(r chi.Router) {
			r.Post("/", handleItem.Create)
			r.Get("/", handleItem.All)
			r.Get("/{id}", handleItem.Get)
			r.Put("/{id}", handleItem.Update)
			r.Delete("/{id}", handleItem.Delete)

			r.Route("/investment", func(r chi.Router) {
				r.Get("/", handleInvestment.All)
				r.Get("/{id}", handleInvestment.Get)
			})

		})
	})

	r.Get("/app.css", handleWebTemplate.Static)

	return r
}
