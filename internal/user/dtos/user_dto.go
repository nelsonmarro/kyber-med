package dtos

import (
	"github.com/nelsonmarro/kyber-med/common/commondtos"
	"github.com/nelsonmarro/kyber-med/internal/user/constanst"
)

type UserDTO struct {
	commondtos.BaseDto
	IDCard string             `json:"idCard"`
	Email  string             `json:"email"`
	Role   constanst.UserRole `json:"role"`
}
