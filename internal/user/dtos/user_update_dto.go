package dtos

import "github.com/nelsonmarro/kyber-med/internal/user/constanst"

type UserUpdateDTO struct {
	IDCard string             `json:"idCard"`
	Email  string             `json:"email"`
	Role   constanst.UserRole `json:"role"`
}
