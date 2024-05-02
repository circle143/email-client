package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"mennr.tech/api/controllers"
)

func Router() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/contact-us", controllers.GetContactUs)
	router.Post("/contact-us", controllers.PostContactUs)

	return router

}
