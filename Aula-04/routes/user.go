package routes

import (
	"github.com/gorilla/mux"
	"github.com/william-cirico/rest-api-golang/handlers"
	"gorm.io/gorm"
)

type userRoutes struct {
	db *gorm.DB
}

func RegisterUserRoutes(router *mux.Router, db *gorm.DB) {
	userRouter := router.PathPrefix("/users").Subrouter()
	userHandler := handlers.NewUserHandler(db)

	userRouter.HandleFunc("/", userHandler.GetUsersHandler).Methods("GET")
	userRouter.HandleFunc("/{id:[0-9]+}", userHandler.GetUserByIdHandler).Methods("GET")
	userRouter.HandleFunc("/", userHandler.CreateUserHandler).Methods("POST")
	userRouter.HandleFunc("/{id:[0-9]+}", userHandler.UpdateUserHandler).Methods("PUT")
}
