package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/wFercho/mines_microservice/internal/application/usecase"
	"github.com/wFercho/mines_microservice/internal/config"
	mongo_db "github.com/wFercho/mines_microservice/internal/db/mongo"
	postgres_db "github.com/wFercho/mines_microservice/internal/db/postgres"
	"github.com/wFercho/mines_microservice/internal/server/handlers"
	"github.com/wFercho/mines_microservice/internal/server/routes"

	"github.com/rs/cors"
)

func main() {
	psconn, _ := postgres_db.ConnectToPostgresDatabase()
	mongodbconn := mongo_db.ConnectToMongoDatabase()

	if config.LoadConfig().Environment == "dev" {
		mine_id := postgres_db.SeedPostgresDB(psconn)
		err := mongo_db.NewMineNodes3DSeeder(mongodbconn).SeedFromCSV("./test_data/mine_nodes3d.csv", mine_id)

		if err != nil {
			fmt.Println("error", err)
		}
	}

	defer mongo_db.DisconnectMongo()

	mineRepo := postgres_db.NewPostgresRepository(psconn)
	mineUsecase := usecase.NewMineUseCase(mineRepo)
	mineHandler := handlers.NewMineHandler(mineUsecase)

	mineNodes3dRepo := mongo_db.NewMineNodes3DMongoRepository(mongodbconn)
	mineNodes3dUseCase := usecase.NewMineNodes3DUseCase(mineNodes3dRepo)
	mineNodes3dHandler := handlers.NewMineNodes3DHandler(mineNodes3dUseCase)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	if config.LoadConfig().Environment == "dev" {
		corsHandler := cors.New(cors.Options{
			AllowedOrigins:   []string{"*"}, // Permite todos los orígenes
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: true,
			MaxAge:           300, // Maximum value not ignored by any of major browsers
		})

		r.Use(corsHandler.Handler)
	}

	routes.RegisterMinesRoutes(r, mineHandler)
	routes.RegisterMineNodes3DRoutes(r, mineNodes3dHandler)

	fmt.Println("Servidor escuchando en :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
