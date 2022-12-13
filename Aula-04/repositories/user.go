package repositories

import (
	"fmt"

	"github.com/william-cirico/rest-api-golang/errors"
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
	Save(user *models.User) (models.User, error)
	Delete(ID uint) error
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (ur *userRepository) Find() (models.Users, error) {
	var users []models.User
	if err := ur.db.Find(&users).Error; err != nil {
		return users, errors.NewHttpError(err, 500, "failed to get users from db")
	}

	return users, nil
}

func (ur *userRepository) FindById(ID uint) (models.User, error) {
	var user models.User
	if err := ur.db.First(&user, ID).Error; err != nil {
		return user, errors.NewHttpError(err, 404, fmt.Sprintf("user with ID: %d not found", ID))
	}

	return user, nil
}

func (ur *userRepository) FindByEmail(email string) (models.User, error) {
	var user models.User
	if err := ur.db.Where("email = ?", email).First(&user).Error; err != nil {
		return user, errors.NewHttpError(err, 404, fmt.Sprintf("user with email: %s not found", email))
	}

	return user, nil
}

func (ur *userRepository) Save(user *models.User) (models.User, error) {
	if err := ur.db.Save(&user).Error; err != nil {
		return *user, errors.NewHttpError(err, 500, "can't save user in db")
	}

	return *user, nil
}

func (ur *userRepository) Delete(ID uint) error {
	if err := ur.db.Delete(&models.User{}, ID).Error; err != nil {
		return errors.NewHttpError(err, 500, "can't delete user in db")
	}

	return nil
}
