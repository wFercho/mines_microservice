package handlers

import (
	"encoding/csv"
	"encoding/json"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/wFercho/mines_microservice/internal/application/usecase"
	"github.com/wFercho/mines_microservice/internal/domain/mine_nodes3d"
	"github.com/wFercho/mines_microservice/internal/server/dto"
)

type MineNodes3DHandler struct {
	UseCase *usecase.MineNodes3DUseCase
}

func NewMineNodes3DHandler(uc *usecase.MineNodes3DUseCase) *MineNodes3DHandler {
	return &MineNodes3DHandler{UseCase: uc}
}

func (h *MineNodes3DHandler) CreateMineNodes3D(w http.ResponseWriter, r *http.Request) {
	var request dto.MineNodes3DRequestDTO

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Error al decodificar el JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	mineNodes, err := request.ToDomain()
	if err != nil {
		http.Error(w, "Error al convertir a entidad del dominio: "+err.Error(), http.StatusInternalServerError)
		return
	}

	createdMineNodes, err := h.UseCase.Create(mineNodes)
	if err != nil {
		http.Error(w, "Error al crear MineNodes3D: "+err.Error(), http.StatusInternalServerError)
		return
	}

	response := dto.FromDomainToMineNodes3Ddto(*createdMineNodes)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (h *MineNodes3DHandler) GetMineNodes3DByID(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		http.Error(w, "Formato de UUID inválido", http.StatusBadRequest)
		return
	}

	mineNodes, err := h.UseCase.FindByID(id)
	if err != nil {
		http.Error(w, "Registro no encontrado: "+err.Error(), http.StatusNotFound)
		return
	}

	response := dto.FromDomainToMineNodes3Ddto(*mineNodes)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *MineNodes3DHandler) GetByMineID(w http.ResponseWriter, r *http.Request) {
	mineIDParam := chi.URLParam(r, "mine_id")
	mineID, err := uuid.Parse(mineIDParam)
	if err != nil {
		http.Error(w, "Formato de UUID inválido", http.StatusBadRequest)
		return
	}

	mineNodes, err := h.UseCase.FindByMineID(mineID)
	if err != nil {
		http.Error(w, "Registro no encontrado: "+err.Error(), http.StatusNotFound)
		return
	}

	response := dto.FromDomainToMineNodes3Ddto(*mineNodes)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *MineNodes3DHandler) DeleteMineNodes3D(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		http.Error(w, "Formato de UUID inválido", http.StatusBadRequest)
		return
	}

	if err := h.UseCase.Delete(id); err != nil {
		http.Error(w, "Error al eliminar: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *MineNodes3DHandler) DeleteByMineID(w http.ResponseWriter, r *http.Request) {
	mineIDParam := chi.URLParam(r, "mine_id")
	mineID, err := uuid.Parse(mineIDParam)
	if err != nil {
		http.Error(w, "Formato de UUID inválido", http.StatusBadRequest)
		return
	}

	if err := h.UseCase.DeleteByMineID(mineID); err != nil {
		http.Error(w, "Error al eliminar registros: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *MineNodes3DHandler) UploadCSV(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error al obtener el archivo: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	nodes, err := parseCSV(file)
	if err != nil {
		http.Error(w, "Error al procesar el CSV: "+err.Error(), http.StatusInternalServerError)
		return
	}

	mineID := uuid.New()

	mineNodes, err := mine_nodes3d.NewMineNodes3D(mineID, nodes)
	if err != nil {
		http.Error(w, "Error al crear la entidad de dominio: "+err.Error(), http.StatusInternalServerError)
		return
	}

	createdMineNodes, err := h.UseCase.Create(mineNodes)
	if err != nil {
		http.Error(w, "Error al guardar los datos: "+err.Error(), http.StatusInternalServerError)
		return
	}

	response := dto.FromDomainToMineNodes3Ddto(*createdMineNodes)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func parseCSV(file multipart.File) ([]mine_nodes3d.Node3D, error) {
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var nodes []mine_nodes3d.Node3D

	for _, record := range records {
		if len(record) < 6 {
			continue
		}

		nodeID := record[0]
		zoneCategory := record[1]
		zoneName := record[2]
		connections := strings.Split(record[3], ";")
		color := record[4]

		x, _ := strconv.Atoi(record[5])
		y, _ := strconv.Atoi(record[6])
		z, _ := strconv.Atoi(record[7])

		node := mine_nodes3d.Node3D{
			ID: nodeID,
			Zone: mine_nodes3d.Zone{
				Category: zoneCategory,
				Name:     zoneName,
			},
			Connections: connections,
			Position: mine_nodes3d.Position{
				X: x,
				Y: y,
				Z: z,
			},
			Color:   color,
			Sensors: []mine_nodes3d.Sensor{},
		}

		nodes = append(nodes, node)
	}

	return nodes, nil
}
