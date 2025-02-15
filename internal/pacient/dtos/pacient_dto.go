package pacient

import (
	"time"

	"github.com/nelsonmarro/kyber-med/common/commondtos"
)

type PacientDto struct {
	commondtos.BaseDto
	FirstName     string    `json:"firstName"`
	LastName      string    `json:"lastName"`
	Email         string    `json:"email"`
	IDCard        string    `json:"idCard"`
	PhoneNumber   string    `json:"phoneNumber"`
	DateOfBirth   time.Time `json:"dateOfBirth"`
	Address       string    `json:"address"`
	Age           uint      `json:"age"`
	Gender        string    `json:"gender"`
	Height        float64   `json:"height"`
	Weight        float64   `json:"weight"`
	TargetWeight  float64   `json:"targetWeight"`
	ActivityLevel string    `json:"activityLevel"`
	DietaryGoal   string    `json:"dietaryGoal"`
	TargetDate    time.Time `json:"targetDate"`
}
