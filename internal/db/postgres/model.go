package db

import (
	"time"

	"github.com/google/uuid"
	"github.com/wFercho/mines_microservice/internal/domain/mine"
)

type MinePostgresModel struct {
	ID                uuid.UUID                   `gorm:"primaryKey"`
	Name              string                      `gorm:"size:100;not null"`
	CompanyName       string                      `gorm:"size:100;not null"`
	Coordinates       GPSCoordinatesPostgresModel `gorm:"embedded;not null"`
	Department        string                      `gorm:"size:100;not null"`
	City              string                      `gorm:"size:100;not null"`
	Neighborhood      string                      `gorm:"size:100;not null"`
	Address           string                      `gorm:"size:255;not null"`
	OperationalStatus string                      `gorm:"size:50;not null"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

type GPSCoordinatesPostgresModel struct {
	Lng float64
	Lat float64
}

func (mm *MinePostgresModel) ToDomain() *mine.Mine {
	return &mine.Mine{
		ID:          mm.ID,
		Name:        mm.Name,
		CompanyName: mm.CompanyName,
		Coordinates: mine.GPSCoordinates{
			Lng: mm.Coordinates.Lng,
			Lat: mm.Coordinates.Lat,
		},
		Department:        mm.Department,
		City:              mm.City,
		Neighborhood:      mm.Neighborhood,
		Address:           mm.Address,
		OperationalStatus: mm.OperationalStatus,
		CreatedAt:         mm.CreatedAt,
		UpdatedAt:         mm.UpdatedAt,
	}
}

func FromDomainToMinePostgresModel(m *mine.Mine) *MinePostgresModel {
	return &MinePostgresModel{
		ID:          m.ID,
		Name:        m.Name,
		CompanyName: m.CompanyName,
		Coordinates: GPSCoordinatesPostgresModel{
			Lng: m.Coordinates.Lng,
			Lat: m.Coordinates.Lat,
		},
		Department:        m.Department,
		City:              m.City,
		Address:           m.Address,
		Neighborhood:      m.Neighborhood,
		OperationalStatus: m.OperationalStatus,
		CreatedAt:         m.CreatedAt,
		UpdatedAt:         m.UpdatedAt,
	}
}
