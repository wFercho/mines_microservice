package models

import "github.com/lib/pq"

type Mine struct {
	ID              uint           `gorm:"primaryKey" json:"id"`
	Nombre          string         `gorm:"size:100;not null" json:"nombre"`
	EmpresaAsociada string         `gorm:"size:100;not null" json:"empresa_asociada"`
	Coordenadas     string         `gorm:"type:geometry(Point,4326);not null" json:"coordenadas"`
	Departamento    string         `gorm:"size:100;not null" json:"departamento"`
	Municipio       string         `gorm:"size:100;not null" json:"municipio"`
	BarrioVereda    string         `gorm:"size:100;not null" json:"barrio_vereda"`
	Direccion       string         `gorm:"size:255;not null" json:"direccion"`
	EstadoOperativo string         `gorm:"size:50;not null" json:"estado_operativo"`
	MQTTTopics      pq.StringArray `gorm:"type:text[]" json:"mqtt_topics"`
}
