package pacient

import (
	"time"

	"github.com/nelsonmarro/kyber-med/common/commonentities"
	uEntities "github.com/nelsonmarro/kyber-med/internal/user/entities"
)

type Pacient struct {
	commonentities.BaseEntity
	UserID      string         `gorm:"type:uuid;not null"`
	User        uEntities.User `gorm:"foreignKey:UserID;references:ID"`
	FirstName   string         `gorm:"type:varchar(100);not null"`
	LastName    string         `gorm:"type:varchar(100);not null"`
	Email       string         `gorm:"type:varchar(100);not null;unique"`
	IDCard      string         `gorm:"type:varchar(15);not null;unique"`
	PhoneNumber string         `gorm:"type:varchar(20)"`
	DateOfBirth time.Time      `gorm:"type:date;not null"`
	Address     string         `gorm:"type:varchar(350)"`
	// Datos básicos
	Age    uint    `gorm:"not null"`
	Gender string  `gorm:"type:varchar(50);not null"`
	Height float64 `gorm:"not null"`
	Weight float64 `gorm:"not null"`

	// Datos de Metas / Objetivos
	TargetWeight  float64    `gorm:"not null"`                   // kg
	ActivityLevel string     `gorm:"type:varchar(100);not null"` // p.ej. "Sedentario", "Ligero", "Moderado", etc.
	DietaryGoal   string     `gorm:"type:varchar(100);not null"` // p.ej. "Perder peso", "Salud mejorada", etc.
	TargetDate    *time.Time `gorm:"type:date"`
}
