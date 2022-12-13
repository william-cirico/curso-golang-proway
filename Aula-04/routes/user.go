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

	userRouter.Handle("", handlers.RootHandler(userHandler.GetUsersHandler)).Methods("GET")
	userRouter.Handle("/{id:[0-9]+}", handlers.RootHandler(userHandler.GetUserByIdHandler)).Methods("GET")
	userRouter.Handle("", handlers.RootHandler(userHandler.CreateUserHandler)).Methods("POST")
	userRouter.Handle("/{id:[0-9]+}", handlers.RootHandler(userHandler.UpdateUserHandler)).Methods("PUT")
	userRouter.Handle("/{id:[0-9]+}", handlers.RootHandler(userHandler.DeleteUserHandler)).Methods("DELETE")
}
