package services

import (
	"github.com/nelsonmarro/kyber-med/common/commonhelpers"
	"github.com/nelsonmarro/kyber-med/internal/user/dtos"
	"github.com/nelsonmarro/kyber-med/internal/user/entities"
	"github.com/nelsonmarro/kyber-med/internal/user/repositories"
)

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (s *userService) RegisterUser(userDto dtos.UserRegisterDto) error {
	usrPwd, err := commonhelpers.GeneratePassword(userDto.Password)
	if err != nil {
		return err
	}

	userDb := entities.User{
		IDCard:   userDto.IDCard,
		Email:    userDto.Email,
		Role:     userDto.Role,
		Password: usrPwd,
	}

	err = s.userRepository.CreateUser(userDb)
	if err != nil {
		return err
	}

	return nil
}
