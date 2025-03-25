package mine

import (
	"time"

	"github.com/google/uuid"
)

type Mine struct {
	ID                uuid.UUID
	Name              string
	CompanyName       string
	Coordinates       GPSCoordinates
	City              string
	Department        string
	Neighborhood      string
	OperationalStatus string
	Address           string

	CreatedAt time.Time
	UpdatedAt time.Time
}

type GPSCoordinates struct {
	Lng float64
	Lat float64
}

func NewMine(name, companyName, deparment, city, neighborhood, address, operationalStatus string, lat, lng float64) (*Mine, error) {
	now := time.Now()
	return &Mine{
		ID:          uuid.New(),
		Name:        name,
		CompanyName: companyName,
		Coordinates: GPSCoordinates{
			Lng: lng,
			Lat: lat,
		},
		Department:        deparment,
		City:              city,
		Neighborhood:      neighborhood,
		OperationalStatus: operationalStatus,
		Address:           address,
		CreatedAt:         now,
		UpdatedAt:         now,
	}, nil
}
