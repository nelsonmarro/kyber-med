package dtos

import "github.com/nelsonmarro/kyber-med/internal/user/constanst"

type UserUpdateDTO struct {
	IDCard string             `json:"idCard"`
	Name   string             `json:"name"`
	Email  string             `json:"email"`
	Role   constanst.UserRole `json:"role"`
}
