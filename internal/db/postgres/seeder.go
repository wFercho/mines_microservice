package db

import (
	"log"

	"github.com/wFercho/mines_microservice/internal/models"
)

func SeedDB() {
	mines := []models.Mine{
		{
			Nombre:          "Mina El Diamante",
			EmpresaAsociada: "Carbones de Boyacá S.A.",
			Coordenadas:     models.GPSCoordinates{Lng: -73.3675, Lat: 5.535},
			Departamento:    "Boyacá",
			Municipio:       "Sogamoso",
			BarrioVereda:    "Vereda Morcá",
			Direccion:       "Km 5 vía Morcá",
			EstadoOperativo: "Activa",
			MQTTTopics:      []string{},
		},
		{
			Nombre:          "Mina La Esperanza",
			EmpresaAsociada: "Minerales del Centro Ltda.",
			Coordenadas:     models.GPSCoordinates{Lng: -73.214, Lat: 5.583},
			Departamento:    "Boyacá",
			Municipio:       "Nobsa",
			BarrioVereda:    "Vereda Las Caleras",
			Direccion:       "Sector Las Caleras",
			EstadoOperativo: "Activa",
			MQTTTopics:      []string{},
		},
	}

	for _, mine := range mines {
		result := POSTGRES_DB.Create(&mine)
		if result.Error != nil {
			log.Printf("❌ Error al insertar mina %s: %v", mine.Nombre, result.Error)
		}
	}

	log.Println("✅ Datos iniciales insertados en la base de datos")
}
