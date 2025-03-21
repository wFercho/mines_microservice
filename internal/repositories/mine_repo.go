package repositories

import (
	"github.com/wFercho/mines_microservice/internal/db"
	"github.com/wFercho/mines_microservice/internal/models"
)

func GetAllMines() ([]models.Mine, error) {
	var mines []models.Mine
	result := db.DB.Find(&mines)
	return mines, result.Error
}
