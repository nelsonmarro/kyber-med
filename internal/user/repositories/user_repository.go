package repositories

import (
	"github.com/nelsonmarro/kyber-med/internal/user/entities"
)

type UserRepository interface {
	GetUserByID(id string) (*entities.User, error)
	GetUserByEmail(email string) (*entities.User, error)
	GetUserByIDCard(idCard string) (*entities.User, error)
	Save(user *entities.User) error
	DeleteUser(user *entities.User) error
}
