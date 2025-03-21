package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/wFercho/mines_microservice/internal/repositories"
)

func GetMinesHandler(w http.ResponseWriter, r *http.Request) {
	mines, err := repositories.GetAllMines()
	if err != nil {
		http.Error(w, "Error al obtener las minas", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mines)
}
