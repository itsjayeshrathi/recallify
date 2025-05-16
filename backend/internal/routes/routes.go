package routes

import (
	"github.com/go-chi/chi"
	"github.com/itsjayeshrathi/recallify/internal/app"
)

func SetupRouts(app *app.Application) *chi.Mux {
	r := chi.NewRouter()
	r.Get("/health", app.HealthCheck)
	return r
}
