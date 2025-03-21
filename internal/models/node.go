package models

import "gorm.io/gorm"

type Node struct {
	ID        string  `gorm:"primaryKey"`
	ZoneType  string  `gorm:"type:varchar(50)"`
	ZoneName  string  `gorm:"type:varchar(100)"`
	PositionX float64 `gorm:"type:float"`
	PositionY float64 `gorm:"type:float"`
	PositionZ float64 `gorm:"type:float"`
	Color     string  `gorm:"type:varchar(20);default:null"`

	Sensors []Sensor `gorm:"foreignKey:NodeID"`
}

func (n *Node) BeforeCreate(tx *gorm.DB) (err error) {
	return nil
}
