package server

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
)

func (s *ChiServer) Init() http.Handler {
	router := s.router
	controllers := s.controllers
	middlewares := s.middlewares

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Recoverer)
	router.Use(render.SetContentType(render.ContentTypeJSON))
	router.Use(middlewares.Logger())
	router.Use(middlewares.HandlePanic())

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   s.config.AllowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           s.config.RouterMaxAge,
	}))

	router.Get("/", controllers.GetHealth)

	router.Group(func(r chi.Router) {
		r.Use(middlewares.Auth())
		r.Use(middlewares.User())

		r.Get("/api/users/{user_id}", controllers.GetUser)
		r.Get("/api/users", controllers.ListUsers)
		r.Post("/api/users", controllers.CreateUser)
		r.Delete("/api/users/{user_id}", controllers.DeleteUser)
		r.Put("/api/users/{user_id}", controllers.ModifyUser)

		r.Get("/api/snap", controllers.GetSnap)
	})

	return router
}
