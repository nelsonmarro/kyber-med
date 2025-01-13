package services

import "github.com/nelsonmarro/kyber-med/internal/user/dtos"

type UserService interface {
	RegisterUser(userDto dtos.UserRegisterDto) error
}
