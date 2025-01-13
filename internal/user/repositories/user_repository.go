package repositories

import "github.com/nelsonmarro/kyber-med/internal/user/entities"

type UserRepository interface {
	CreateUser(entities.User) error
}
