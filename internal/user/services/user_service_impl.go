package user

import (
	"github.com/nelsonmarro/kyber-med/common/commondtos"
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

func (s *userService) GetUserById(id string) (*dtos.UserDTO, error) {
	dbUser, error := s.userRepository.GetUserByID(id)
	if error != nil {
		return nil, error
	}

	return &dtos.UserDTO{
		BaseDto: commondtos.BaseDto{
			ID:        dbUser.ID,
			CreatedAt: dbUser.CreatedAt,
		},
		IDCard: dbUser.IDCard,
		Email:  dbUser.Email,
		Role:   dbUser.Role,
	}, nil
}

func (s *userService) GetUserWithPasswordByEmail(email string) (*dtos.UserDTO, string, error) {
	dbUser, error := s.userRepository.GetUserByEmail(email)
	if error != nil {
		return nil, "", error
	}

	return &dtos.UserDTO{
		BaseDto: commondtos.BaseDto{
			ID:        dbUser.ID,
			CreatedAt: dbUser.CreatedAt,
		},
		Name:   dbUser.Name,
		IDCard: dbUser.IDCard,
		Email:  dbUser.Email,
		Role:   dbUser.Role,
	}, dbUser.Password, nil
}

func (s *userService) GetUserWithPassswordByIDCard(idCard string) (*dtos.UserDTO, string, error) {
	dbUser, error := s.userRepository.GetUserByIDCard(idCard)
	if error != nil {
		return nil, "", error
	}

	return &dtos.UserDTO{
		BaseDto: commondtos.BaseDto{
			ID:        dbUser.ID,
			CreatedAt: dbUser.CreatedAt,
		},
		IDCard: dbUser.IDCard,
		Email:  dbUser.Email,
		Role:   dbUser.Role,
	}, dbUser.Password, nil
}

func (s *userService) ValidUser(id string, password string) bool {
	user, _ := s.userRepository.GetUserByID(id)
	if user.Email == "" || user.IDCard == "" {
		return false
	}

	return commonhelpers.CheckPasswordHash(user.Password, password)
}

func (s *userService) RegisterUser(userDto dtos.UserRegisterDTO) error {
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

	err = s.userRepository.Save(&userDb)
	if err != nil {
		return err
	}

	return nil
}

func (s *userService) UpdateUser(userDto dtos.UserUpdateDTO, id string) error {
	dbUser, err := s.userRepository.GetUserByID(id)
	if err != nil {
		return err
	}

	dbUser.IDCard = userDto.IDCard
	dbUser.Email = userDto.Email
	dbUser.Role = userDto.Role

	err = s.userRepository.Save(dbUser)
	if err != nil {
		return err
	}

	return nil
}

func (s *userService) DeleteUser(id string) error {
	dbUser, err := s.userRepository.GetUserByID(id)
	if err != nil {
		return err
	}

	err = s.userRepository.DeleteUser(dbUser)
	if err != nil {
		return err
	}

	return nil
}
