package db

import (
	"log"

	"github.com/wFercho/mines_microservice/internal/models"
)

func SeedDB() {
	mines := []models.Mine{
		{Nombre: "Mina El Diamante", EmpresaAsociada: "Carbones de Boyacá S.A.", Coordenadas: "POINT(-73.3675 5.535)", Departamento: "Boyacá", Municipio: "Sogamoso", BarrioVereda: "Vereda Morcá", Direccion: "Km 5 vía Morcá", EstadoOperativo: "Activa", MQTTTopics: []string{}},
		{Nombre: "Mina La Esperanza", EmpresaAsociada: "Minerales del Centro Ltda.", Coordenadas: "POINT(-73.214 5.583)", Departamento: "Boyacá", Municipio: "Nobsa", BarrioVereda: "Vereda Las Caleras", Direccion: "Sector Las Caleras", EstadoOperativo: "Activa", MQTTTopics: []string{}},
	}

	for _, mine := range mines {
		DB.Create(&mine)
	}

	log.Println("✅ Datos iniciales insertados en la base de datos")
}
