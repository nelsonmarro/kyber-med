package pacient

import (
	"time"

	"github.com/nelsonmarro/kyber-med/common/commondtos"
)

type PacientDto struct {
	commondtos.BaseDto
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	Email       string    `json:"email"`
	IDCard      string    `json:"idCard"`
	PhoneNumber string    `json:"phoneNumber"`
	DateOfBirth time.Time `json:"dateOfBirth"`
	Gender      string    `json:"gender"`
	Address     string    `json:"address"`
	Age         uint      `gorm:"not null"`
	Gender      string    `gorm:"type:varchar(50);not null"`
	Height      float64   `gorm:"not null"`
	Weight      float64   `gorm:"not null"`

	// Datos de Metas / Objetivos
	TargetWeight  float64    `gorm:"not null"`                   // kg
	ActivityLevel string     `gorm:"type:varchar(100);not null"` // p.ej. "Sedentario", "Ligero", "Moderado", etc.
	DietaryGoal   string     `gorm:"type:varchar(100);not null"` // p.ej. "Perder peso", "Salud mejorada", etc.
	TargetDate    *time.Time `gorm:"type:date"`
}
