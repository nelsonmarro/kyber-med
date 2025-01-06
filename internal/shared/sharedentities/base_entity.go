package sharedentities

import (
	"gorm.io/gorm"
)

type BaseEntity struct {
	ID        uint `gorm:"primaryKey;autoIncrement"`
	CreatedAt int
	UpdatedAt int
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
