package handlers

import (
	"net/http"

	"github.com/william-cirico/rest-api-golang/usecases"
	"gorm.io/gorm"
)

type userHandler struct {
	db *gorm.DB
	uc usecases.UserUseCase
}

type UserHandler interface {
	GetUsersHandler(w http.ResponseWriter, r *http.Request)
	GetUserByIdHandler(w http.ResponseWriter, r *http.Request)
	CreateUserHandler(w http.ResponseWriter, r *http.Request)
	UpdateUserHandler(w http.ResponseWriter, r *http.Request)
	DeleteUserHandler(w http.ResponseWriter, r *http.Request)
}

func NewUserHandler(db *gorm.DB) UserHandler {
	uc := usecases.NewUserUseCase(db)

	return &userHandler{
		db: db,
		uc: uc,
	}
}

func (uh *userHandler) GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := uh.uc.GetUsers()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, users.ToJSON())
}

func (uh *userHandler) GetUserByIdHandler(w http.ResponseWriter, r *http.Request) {}

func (uh *userHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {}

func (uh *userHandler) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {}

func (uh *userHandler) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {}
