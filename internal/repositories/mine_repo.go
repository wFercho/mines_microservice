package repositories

import (
	db "github.com/wFercho/mines_microservice/internal/db/postgres"
	"github.com/wFercho/mines_microservice/internal/models"
)

func GetAllMines() ([]models.Mine, error) {
	var mines []models.Mine
	result := db.POSTGRES_DB.Find(&mines)
	return mines, result.Error
}
