package routes

import (
	"github.com/KyTama/KyGo/app"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	//newrelic "github.com/newrelic/go-agent"
)

func NewRoutes() http.Handler {
	r := chi.NewRouter()
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-Requested-With", "access-token", "initial-code"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
	})
	r.Use(cors.Handler)
	r.Use(middleware.Throttle(100))
	r.Use(middleware.Logger)
	registerRoutes(r)
	return http.TimeoutHandler(r, 30*time.Second, `{"Message": "Service Unavailable"}`)
}

func registerRoutes(r *chi.Mux) *chi.Mux {
	r.Group(func(r chi.Router) {
		r.Get("/ping", app.Main)
	})

	return r
}
