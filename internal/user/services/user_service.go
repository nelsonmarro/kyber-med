package services

import "github.com/nelsonmarro/kyber-med/internal/user/dtos"

type UserService interface {
	RegisterUser(userDto dtos.UserRegisterDTO) error
	ValidUser(id, password string) bool
	GetUserById(id string) (*dtos.UserDTO, error)
	GetUserWithPasswordByEmail(email string) (*dtos.UserDTO, string, error)
	GetUserWithPassswordByIDCard(idCard string) (*dtos.UserDTO, string, error)
	UpdateUser(userDto dtos.UserUpdateDTO, id string) error
	DeleteUser(id string) error
}
