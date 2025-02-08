package pacient

import "time"

type UpsertPacientDto struct {
	FirstName             string    `json:"firstName"`
	LastName              string    `json:"lastName"`
	Email                 string    `json:"email"`
	IDCard                string    `json:"idCard"`
	PhoneNumber           string    `json:"phoneNumber,omitempty"`
	DateOfBirth           time.Time `json:"dateOfBirth"`
	Gender                string    `json:"gender,omitempty"`
	Address               string    `json:"address,omitempty"`
	EmergencyContactName  string    `json:"emergencyContactName,omitempty"`
	EmergencyContactPhone string    `json:"emergencyContactPhone,omitempty"`
}
