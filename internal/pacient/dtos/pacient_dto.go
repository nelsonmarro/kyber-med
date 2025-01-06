package dtos

import (
	"time"

	"github.com/nelsonmarro/kyber-med/internal/shared/shareddtos"
)

type PacientDto struct {
	shareddtos.BaseDto
	FirstName             string    `json:"firstName"`
	LastName              string    `json:"lastName"`
	Email                 string    `json:"email"`
	IDCard                string    `json:"idCard"`
	PhoneNumber           string    `json:"phoneNumber"`
	DateOfBirth           time.Time `json:"dateOfBirth"`
	Gender                string    `json:"gender"`
	Address               string    `json:"address"`
	EmergencyContactName  string    `json:"emergencyContactName"`
	EmergencyContactPhone string    `json:"emergencyContactPhone"`
}
