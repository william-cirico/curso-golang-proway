package infrastructure

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/william-cirico/rest-api-golang/routes"
	"gorm.io/gorm"
)

type muxRouter struct {
	db *gorm.DB
}

func NewMuxRouter(db *gorm.DB) Router {
	return &muxRouter{db: db}
}

func (mr *muxRouter) Start() {
	router := mux.NewRouter()

	// Aqui será necessário registrar todas as rotas de usuários
	routes.RegisterUserRoutes(router, mr.db)

	// Adicionando os middlewares de log e cors
	handler := handlers.LoggingHandler(os.Stdout, handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "PATH", "DELETE", "OPTIONS"}),
	)(router))

	srv := &http.Server{
		Handler: handler,
		Addr:    "127.0.0.1:" + os.Getenv("PORT"),
	}

	fmt.Println("Server listening on port: " + os.Getenv("PORT"))
	log.Fatal(srv.ListenAndServe())
}
