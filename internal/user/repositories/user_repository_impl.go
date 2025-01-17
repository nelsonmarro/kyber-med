package repositories

import (
	"errors"

	"github.com/nelsonmarro/kyber-med/internal/database"
	"github.com/nelsonmarro/kyber-med/internal/user/entities"
	"gorm.io/gorm"
)

type userRepositoryImpl struct {
	db database.Database
}

func NewUserRepository(db database.Database) UserRepository {
	return &userRepositoryImpl{db: db}
}

func (r *userRepositoryImpl) GetUserByEmail(email string) (*entities.User, error) {
	db := r.db.GetDb()

	var user entities.User
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepositoryImpl) GetUserByIDCard(idCard string) (*entities.User, error) {
	db := r.db.GetDb()

	var user entities.User
	if err := db.Where(&entities.User{Email: idCard}).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepositoryImpl) GetUserByID(id string) (*entities.User, error) {
	db := r.db.GetDb()

	var user entities.User
	result := db.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (r *userRepositoryImpl) Save(user *entities.User) error {
	db := r.db.GetDb()
	result := db.Save(user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *userRepositoryImpl) DeleteUser(user *entities.User) error {
	db := r.db.GetDb()

	result := db.Delete(user)
	if err := result.Error; err != nil {
		return err
	}

	return nil
}
