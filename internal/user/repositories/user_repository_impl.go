package repositories

import (
	"github.com/nelsonmarro/kyber-med/internal/database"
	uEntities "github.com/nelsonmarro/kyber-med/internal/user/entities"
)

type userRepositoryImpl struct {
	db database.Database
}

func NewUserRepository(db database.Database) UserRepository {
	return &userRepositoryImpl{db: db}
}

func (r *userRepositoryImpl) GetUserByEmail(email string) (*uEntities.User, error) {
	db := r.db.GetDb()

	var user uEntities.User
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepositoryImpl) GetUserByIDCard(idCard string) (*uEntities.User, error) {
	db := r.db.GetDb()

	var user uEntities.User
	if err := db.Where(&uEntities.User{Email: idCard}).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepositoryImpl) GetUserByID(id string) (*uEntities.User, error) {
	db := r.db.GetDb()

	var user uEntities.User
	result := db.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (r *userRepositoryImpl) Save(user *uEntities.User) error {
	db := r.db.GetDb()
	result := db.Save(user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *userRepositoryImpl) DeleteUser(user *uEntities.User) error {
	db := r.db.GetDb()

	result := db.Delete(user)
	if err := result.Error; err != nil {
		return err
	}

	return nil
}
