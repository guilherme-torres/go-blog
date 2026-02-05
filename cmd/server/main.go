package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	app_errors "github.com/guilherme-torres/go-blog/internal/errors"
	"github.com/guilherme-torres/go-blog/internal/handlers"
	"github.com/guilherme-torres/go-blog/internal/middlewares"
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

	userRepo := repositories.NewUserRepo(db)
	articleRepo := repositories.NewArticleRepo(db)
	
	userService := services.NewUserService(userRepo)
	authService := services.NewAuthService(userRepo, redisClient)
	articleService := services.NewArticleService(articleRepo)
	
	userHandler := handlers.NewUserHandler(userService, articleService)
	authHandler := handlers.NewAuthHandler(authService)
	articleHandler := handlers.NewArticleHandler(articleService)

	fs := http.FileServer(http.Dir("assets"))
	mux := http.NewServeMux()

	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	mux.HandleFunc("POST /admin/login", app_errors.HandleErrors(authHandler.Login))
	mux.HandleFunc("GET /admin/login", app_errors.HandleErrors(authHandler.Login))
	mux.HandleFunc("POST /admin/logout", app_errors.HandleErrors(middlewares.AuthMiddleware(authHandler.Logout, authService)))
	mux.HandleFunc("GET /admin", app_errors.HandleErrors(
		middlewares.AuthMiddleware(
			middlewares.VerifyRole(userHandler.Admin, []string{"admin", "editor"}, authService),
			authService,
		),
	))
	mux.HandleFunc("POST /admin/users/new", app_errors.HandleErrors(
		middlewares.AuthMiddleware(
			middlewares.VerifyRole(userHandler.CreateUser, []string{"admin"}, authService),
			authService,
		),
	))
	mux.HandleFunc("GET /admin/users", app_errors.HandleErrors(
		middlewares.AuthMiddleware(
			middlewares.VerifyRole(userHandler.ListUsers, []string{"admin"}, authService),
			authService,
		),
	))
	mux.HandleFunc("GET /admin/users/{id}/delete", app_errors.HandleErrors(
		middlewares.AuthMiddleware(
			middlewares.VerifyRole(userHandler.DeleteUser, []string{"admin"}, authService),
			authService,
		),
	))
	mux.HandleFunc("POST /admin/articles/new", app_errors.HandleErrors(
		middlewares.AuthMiddleware(
			middlewares.VerifyRole(articleHandler.CreateArticle, []string{"admin", "editor"}, authService),
			authService,
		),
	))
	mux.HandleFunc("GET /admin/articles/{id}/delete", app_errors.HandleErrors(
		middlewares.AuthMiddleware(
			middlewares.VerifyRole(articleHandler.DeleteArticle, []string{"admin", "editor"}, authService),
			authService,
		),
	))
	mux.HandleFunc("GET /articles/{id}", app_errors.HandleErrors(articleHandler.GetArticle))
	mux.HandleFunc("GET /articles", app_errors.HandleErrors(articleHandler.ListArticles))

	if err := http.ListenAndServe(":8000", mux); err != nil {
		log.Fatal("Erro ao iniciar o servidor:", err)
	}
}
