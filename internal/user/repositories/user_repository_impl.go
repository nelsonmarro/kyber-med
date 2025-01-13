package repositories

import (
	"github.com/nelsonmarro/kyber-med/internal/database"
	"github.com/nelsonmarro/kyber-med/internal/user/entities"
)

type userRepositoryImpl struct {
	db database.Database
}

func NewUserRepository(db database.Database) UserRepository {
	return &userRepositoryImpl{db: db}
}

func (r userRepositoryImpl) CreateUser(user entities.User) error {
	db := r.db.GetDb()

	res := db.Create(&user)
	if res.Error != nil {
		return res.Error
	}

	return nil
}
