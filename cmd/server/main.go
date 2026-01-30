package main

import (
	"context"
	"database/sql"
	"html/template"
	"log"
	"net/http"

	app_errors "github.com/guilherme-torres/go-blog/internal/errors"
	"github.com/guilherme-torres/go-blog/internal/handlers"
	"github.com/guilherme-torres/go-blog/internal/repositories"
	"github.com/guilherme-torres/go-blog/internal/services"
	"github.com/guilherme-torres/go-blog/internal/utils"
	_ "github.com/mattn/go-sqlite3"
	"github.com/redis/go-redis/v9"
)

func main() {
	db, err := sql.Open("sqlite3", "./blog.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		log.Fatal("Não foi possível conectar ao banco de dados:", err)
	}
	userRepo := repositories.NewUserRepo(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
		Protocol: 2,
	})
	if err := rdb.Ping(ctx).Err(); err != nil {
		log.Fatal("Não foi possível conectar ao redis:", err)
	}
	redisClient := utils.NewRedisClient(rdb)
	authService := services.NewAuthService(userRepo, redisClient)
	authHandler := handlers.NewAuthHandler(authService)

	mux := http.NewServeMux()

	mux.HandleFunc("POST /users", app_errors.HandleErrors(userHandler.CreateUser))
	mux.HandleFunc("POST /auth/login", app_errors.HandleErrors(authHandler.Login))
	mux.HandleFunc("GET /auth/login", app_errors.HandleErrors(authHandler.Login))
	mux.HandleFunc("GET /admin", app_errors.HandleErrors(func(w http.ResponseWriter, r *http.Request) error {
		tmpl := template.Must(template.ParseFiles("./assets/templates/admin.html"))
		tmpl.Execute(w, nil)
		return nil
	}))

	if err := http.ListenAndServe(":8000", mux); err != nil {
		log.Fatal("Erro ao iniciar o servidor:", err)
	}
}
