package models

import "gorm.io/gorm"

type SensorData struct {
	ID         uint   `gorm:"primaryKey;autoIncrement"`
	SensorID   string `gorm:"type:varchar(50);index"`
	Type       string `gorm:"type:varchar(50)"`
	Value      float64
	Unit       string `gorm:"type:varchar(20)"`
	AlertName  string `gorm:"type:varchar(50)"`
	AlertColor string `gorm:"type:varchar(20)"`

	Sensor Sensor `gorm:"foreignKey:SensorID"`
}

func (sd *SensorData) BeforeCreate(tx *gorm.DB) (err error) {
	// Aquí podrías implementar lógica de validación, por ejemplo, para las alertas.
	return nil
}
