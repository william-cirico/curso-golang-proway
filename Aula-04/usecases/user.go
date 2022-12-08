package usecases

import (
	"errors"

	"github.com/william-cirico/rest-api-golang/models"
	"github.com/william-cirico/rest-api-golang/repositories"
	"gorm.io/gorm"
)

type userUseCase struct {
	db *gorm.DB
	r  repositories.UserRepository
}

type UserUseCase interface {
	GetUsers() (models.Users, error)
	GetUserById(ID uint) (models.User, error)
	CreateUser(user models.User) (models.User, error)
	UpdateUser(user models.User) (models.User, error)
	DeleteUser(ID uint) error
}

func NewUserUseCase(db *gorm.DB) UserUseCase {
	ur := repositories.NewUserRepository(db)

	return &userUseCase{db: db, r: ur}
}

func (uc *userUseCase) GetUsers() (models.Users, error) {
	var users []models.User

	return users, errors.New("not implemented yet")
}

func (uc *userUseCase) GetUserById(ID uint) (models.User, error) {
	return models.User{}, errors.New("not implemented yet")
}

func (uc *userUseCase) CreateUser(user models.User) (models.User, error) {
	return models.User{}, errors.New("not implemented yet")
}

func (uc *userUseCase) UpdateUser(user models.User) (models.User, error) {
	return models.User{}, errors.New("not implemented yet")
}

func (uc *userUseCase) DeleteUser(ID uint) error {
	return errors.New("not implemented yet")
}
