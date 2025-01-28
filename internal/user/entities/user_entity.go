package entities

import (
	"github.com/nelsonmarro/kyber-med/common/commonentities"
	"github.com/nelsonmarro/kyber-med/internal/user/constanst"
)

type User struct {
	commonentities.BaseEntity
	IDCard   string             `gorm:"type:varchar(10);unique;not null"`
	Name     string             `gorm:"type:varchar(100);not null"`
	Email    string             `gorm:"type:varchar(200);unique;not null"`
	Password string             `gorm:"not null"`
	Role     constanst.UserRole `gorm:"type:varchar(10);default:'user'"`
}
