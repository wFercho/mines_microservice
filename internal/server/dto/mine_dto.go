package dto

import "github.com/wFercho/mines_microservice/internal/domain/mine"

type MineRequestDTO struct {
	Name              string  `json:"name" validate:"required"`
	CompanyName       string  `json:"company_name" validate:"required"`
	Department        string  `json:"department" validate:"required"`
	City              string  `json:"city" validate:"required"`
	Neighborhood      string  `json:"neighborhood"`
	Address           string  `json:"address" validate:"required"`
	OperationalStatus string  `json:"operational_status"`
	Lat               float64 `json:"lat" validate:"required"`
	Lng               float64 `json:"lng" validate:"required"`
}

type MineResponseDTO struct {
	ID                string  `json:"id"`
	Name              string  `json:"name"`
	CompanyName       string  `json:"company_name"`
	Department        string  `json:"department"`
	City              string  `json:"city"`
	Neighborhood      string  `json:"neighborhood"`
	Address           string  `json:"address"`
	OperationalStatus string  `json:"operational_status"`
	Lat               float64 `json:"lat"`
	Lng               float64 `json:"lng"`
}

func (dto *MineRequestDTO) ToDomain() (*mine.Mine, error) {
	return mine.NewMine(
		dto.Name,
		dto.CompanyName,
		dto.Department,
		dto.City,
		dto.Neighborhood,
		dto.Address,
		dto.OperationalStatus,
		dto.Lat,
		dto.Lng,
	)
}

func FromMineDomain(mine mine.Mine) MineResponseDTO {
	return MineResponseDTO{
		ID:                mine.ID.String(),
		Name:              mine.Name,
		CompanyName:       mine.CompanyName,
		Department:        mine.Department,
		City:              mine.City,
		Neighborhood:      mine.Neighborhood,
		Address:           mine.Address,
		OperationalStatus: mine.OperationalStatus,
		Lat:               mine.Coordinates.Lat,
		Lng:               mine.Coordinates.Lng,
	}
}
