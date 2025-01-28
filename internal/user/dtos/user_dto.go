package dtos

import (
	"github.com/nelsonmarro/kyber-med/common/commondtos"
	"github.com/nelsonmarro/kyber-med/internal/user/constanst"
)

type UserDTO struct {
	commondtos.BaseDto
	IDCard string             `json:"idCard"`
	Name   string             `json:"name"`
	Email  string             `json:"email"`
	Role   constanst.UserRole `json:"role"`
}
