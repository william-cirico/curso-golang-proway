package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/william-cirico/rest-api-golang/errors"
	"github.com/william-cirico/rest-api-golang/models"
	"github.com/william-cirico/rest-api-golang/schemas"
	"github.com/william-cirico/rest-api-golang/usecases"
	"gorm.io/gorm"
)

type userHandler struct {
	db *gorm.DB
	uc usecases.UserUseCase
}

type UserHandler interface {
	GetUsersHandler(w http.ResponseWriter, r *http.Request) error
	GetUserByIdHandler(w http.ResponseWriter, r *http.Request) error
	CreateUserHandler(w http.ResponseWriter, r *http.Request) error
	UpdateUserHandler(w http.ResponseWriter, r *http.Request) error
	DeleteUserHandler(w http.ResponseWriter, r *http.Request) error
}

func NewUserHandler(db *gorm.DB) UserHandler {
	uc := usecases.NewUserUseCase(db)

	return &userHandler{
		db: db,
		uc: uc,
	}
}

func (uh *userHandler) GetUsersHandler(w http.ResponseWriter, r *http.Request) error {
	users, err := uh.uc.GetUsers()
	if err != nil {
		return err
	}

	return respondWithJSON(w, http.StatusOK, users.ToJSON())
}

func (uh *userHandler) GetUserByIdHandler(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]

	ID, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	user, err := uh.uc.GetUserById(uint(ID))

	if err != nil {
		return err
	}

	return respondWithJSON(w, http.StatusOK, user)
}

func (uh *userHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) error {
	var userSchema schemas.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&userSchema); err != nil {
		return errors.NewHttpError(err, 400, "bad parameters")
	}

	user := models.User{
		Name:     userSchema.Name,
		Email:    userSchema.Email,
		Password: userSchema.Password,
	}

	user, err := uh.uc.CreateUser(user)
	if err != nil {
		return err
	}

	return respondWithJSON(w, http.StatusCreated, user)
}

func (uh *userHandler) UpdateUserHandler(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]

	ID, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	user, err := uh.uc.GetUserById(uint(ID))
	if err != nil {
		return err
	}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		return errors.NewHttpError(err, 400, "bad parameters")
	}

	user, err = uh.uc.UpdateUser(user)
	if err != nil {
		return err
	}

	return respondWithJSON(w, http.StatusOK, user)
}

func (uh *userHandler) DeleteUserHandler(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]

	ID, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	err = uh.uc.DeleteUser(uint(ID))
	if err != nil {
		return err
	}

	return respondWithJSON(w, http.StatusNoContent, nil)
}
