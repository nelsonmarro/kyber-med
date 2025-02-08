package user

import uDtos "github.com/nelsonmarro/kyber-med/internal/user/dtos"

type UserService interface {
	RegisterUser(userDto uDtos.UserRegisterDTO) error
	ValidUser(id, password string) bool
	GetUserById(id string) (*uDtos.UserDTO, error)
	GetUserWithPasswordByEmail(email string) (*uDtos.UserDTO, string, error)
	GetUserWithPassswordByIDCard(idCard string) (*uDtos.UserDTO, string, error)
	UpdateUser(userDto uDtos.UserUpdateDTO, id string) error
	DeleteUser(id string) error
}
