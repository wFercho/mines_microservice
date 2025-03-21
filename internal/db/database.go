package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/wFercho/mines_microservice/internal/config"
)

func ConnectDatabase() (*gorm.DB, error) {
	cfg := config.LoadConfig()
	dsn := cfg.GetPostgresDatabaseURL()

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error al conectar con la base de datos: %v", err)
		return nil, err
	}

	return db, nil
}
