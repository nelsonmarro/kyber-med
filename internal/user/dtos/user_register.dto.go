package dtos

import "github.com/nelsonmarro/kyber-med/internal/user/constanst"

type UserRegisterDto struct {
	IDCard   string             `gorm:"type:varchar(10);unique;not null"`
	Email    string             `gorm:"type:varchar(200);unique;not null"`
	Password string             `gorm:"not null"`
	Role     constanst.UserRole `gorm:"type:varchar(10);default:'user'"`
}
