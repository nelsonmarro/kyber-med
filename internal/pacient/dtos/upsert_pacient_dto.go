package pacient

import "time"
	ActivityLevel "github.com/nelsonmarro/kyber-med/internal/pacient/enums"

type UpsertPacientDto struct {
	FirstName     string    `json:"firstName"`
	LastName      string    `json:"lastName"`
	Email         string    `json:"email"`
	IDCard        string    `json:"idCard"`
	PhoneNumber   string    `json:"phoneNumber,omitempty"`
	DateOfBirth   time.Time `json:"dateOfBirth"`
	Address       string    `json:"address,omitempty"`
	Age           uint      `json:"age"`
	Gender        string    `json:"gender"`
	Height        float64   `json:"height"`
	Weight        float64   `json:"weight"`
	TargetWeight  float64   `json:"targetWeight"`
	ActivityLevel A    `json:"activityLevel"`
	DietaryGoal   string    `json:"dietaryGoal"`
	TargetDate    time.Time `json:"targetDate,omitempty"`
}
