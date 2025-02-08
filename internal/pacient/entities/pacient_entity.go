package pacient

import (
	"time"

	"github.com/nelsonmarro/kyber-med/common/commonentities"
	uEntities "github.com/nelsonmarro/kyber-med/internal/user/entities"
)

type Pacient struct {
	commonentities.BaseEntity
	FirstName             string         `gorm:"type:varchar(100);not null"`
	LastName              string         `gorm:"type:varchar(100);not null"`
	Email                 string         `gorm:"type:varchar(100);not null;unique"`
	IDCard                string         `gorm:"type:varchar(15);not null;uniqueIndex"`
	PhoneNumber           string         `gorm:"type:varchar(20)"`
	DateOfBirth           time.Time      `gorm:"type:date;not null"`
	Gender                string         `gorm:"type:varchar(10)"`
	Address               string         `gorm:"type:varchar(350)"`
	EmergencyContactName  string         `gorm:"type:varchar(100)"`
	EmergencyContactPhone string         `gorm:"type:varchar(15)"`
	User                  uEntities.User `gorm:"foreignKey:ID"`
}
