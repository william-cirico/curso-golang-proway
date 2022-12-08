package repositories

import (
	"errors"

	"github.com/william-cirico/rest-api-golang/models"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

type UserRepository interface {
	Find() (models.Users, error)
	FindById(ID uint) (models.User, error)
	FindByEmail(email string) (models.User, error)
	Save(models.User) (models.User, error)
	Delete(ID uint) error
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (ur *userRepository) Find() (models.Users, error) {
	var users []models.User

	return users, errors.New("not implemented")
}

func (ur *userRepository) FindById(ID uint) (models.User, error) {
	return models.User{}, errors.New("not implemented")
}

func (ur *userRepository) FindByEmail(email string) (models.User, error) {
	return models.User{}, errors.New("not implemented")
}

func (ur *userRepository) Save(models.User) (models.User, error) {
	return models.User{}, errors.New("not implemented")
}

func (ur *userRepository) Delete(ID uint) error {
	return errors.New("not implemented")
}
