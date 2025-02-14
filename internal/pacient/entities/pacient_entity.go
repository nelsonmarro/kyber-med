package pacient

import (
	"time"

	"github.com/nelsonmarro/kyber-med/common/commonentities"
	uEntities "github.com/nelsonmarro/kyber-med/internal/user/entities"
)

type Pacient struct {
	commonentities.BaseEntity
	FirstName   string         `gorm:"type:varchar(100);not null"`
	LastName    string         `gorm:"type:varchar(100);not null"`
	Email       string         `gorm:"type:varchar(100);not null;unique"`
	IDCard      string         `gorm:"type:varchar(15);not null;unique"`
	PhoneNumber string         `gorm:"type:varchar(20)"`
	DateOfBirth time.Time      `gorm:"type:date;not null"`
	Gender      string         `gorm:"type:varchar(10)"`
	Address     string         `gorm:"type:varchar(350)"`
	UserID      string         `gorm:"type:uuid;not null"`
	User        uEntities.User `gorm:"foreignKey:UserID;references:ID"`
	// Datos básicos
	Edad   uint    `gorm:"not null"`
	Genero string  `gorm:"type:varchar(50);not null"`
	Altura float64 `gorm:"not null"`
	Peso   float64 `gorm:"not null"`

	// Datos de Metas / Objetivos
	PesoObjetivo      float64    `gorm:"not null"`                   // kg
	NivelActividad    string     `gorm:"type:varchar(100);not null"` // p.ej. "Sedentario", "Ligero", "Moderado", etc.
	ObjetivoDietetico string     `gorm:"type:varchar(100);not null"` // p.ej. "Perder peso", "Salud mejorada", etc.
	FechaMeta         *time.Time `gorm:"type:date"`
}
