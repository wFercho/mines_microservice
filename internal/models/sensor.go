package models

import "gorm.io/gorm"

type Sensor struct {
	ID     string `gorm:"primaryKey"`
	Type   string `gorm:"type:varchar(50)"`
	Unit   string `gorm:"type:varchar(20)"`
	NodeID string `gorm:"type:varchar(50);index"`

	Node Node `gorm:"foreignKey:NodeID"`
}

func (s *Sensor) BeforeCreate(tx *gorm.DB) (err error) {
	// Puedes agregar validaciones o inicializaciones antes de guardar
	return nil
}
