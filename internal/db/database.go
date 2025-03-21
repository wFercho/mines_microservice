package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/wFercho/mines_microservice/internal/config"
	"github.com/wFercho/mines_microservice/internal/models"
)

var DB *gorm.DB

func ConnectDatabase() (*gorm.DB, error) {
	cfg := config.LoadConfig()
	dsn := cfg.GetPostgresDatabaseURL()

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error al conectar con la base de datos: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&models.Mine{})
	if err != nil {
		log.Fatal("❌ Error en la migración:", err)
	}

	DB = db
	return db, nil
}
