package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/wFercho/mines_microservice/internal/config"
	"github.com/wFercho/mines_microservice/internal/db"
	"github.com/wFercho/mines_microservice/internal/routes"
)

func main() {
	db.ConnectToPostgresDatabase()

	if config.LoadConfig().Environment == "dev" {
		db.SeedDB()
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Agregar rutas
	r.Mount("/api", routes.MineRoutes())

	fmt.Println("Servidor escuchando en :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
