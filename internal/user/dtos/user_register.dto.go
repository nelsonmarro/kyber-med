package dtos

import "github.com/nelsonmarro/kyber-med/internal/user/constanst"

type UserRegisterDTO struct {
	IDCard   string             `json:"idCard"`
	Email    string             `json:"email"`
	Name     string             `json:"name"`
	Password string             `json:"password"`
	Role     constanst.UserRole `json:"role"`
}
