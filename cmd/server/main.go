package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/guilherme-torres/go-blog/internal/errors"
	"github.com/guilherme-torres/go-blog/internal/handlers"
	"github.com/guilherme-torres/go-blog/internal/repositories"
	"github.com/guilherme-torres/go-blog/internal/services"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./blog.db")
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal("Não foi possível conectar ao banco:", err)
	}
	defer db.Close()

	userRepo := repositories.NewUserRepo(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	mux := http.NewServeMux()

	mux.HandleFunc("POST /users", app_errors.HandleErrors(userHandler.CreateUser))

	http.ListenAndServe(":8000", mux)
}
