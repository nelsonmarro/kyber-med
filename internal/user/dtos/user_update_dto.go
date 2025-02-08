package user

import uConstanst "github.com/nelsonmarro/kyber-med/internal/user/constanst"

type UserUpdateDTO struct {
	IDCard string              `json:"idCard"`
	Name   string              `json:"name"`
	Email  string              `json:"email"`
	Role   uConstanst.UserRole `json:"role"`
}
