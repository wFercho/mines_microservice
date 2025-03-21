package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/wFercho/mines_microservice/internal/handlers"
)

func MineRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/mines", handlers.GetMinesHandler)
	return r
}
