package repositories

import (
	uEntities "github.com/nelsonmarro/kyber-med/internal/user/entities"
)

type UserRepository interface {
	GetUserByID(id string) (*uEntities.User, error)
	GetUserByEmail(email string) (*uEntities.User, error)
	GetUserByIDCard(idCard string) (*uEntities.User, error)
	Save(user *uEntities.User) error
	DeleteUser(user *uEntities.User) error
}
