package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"

	"github.com/guilherme-torres/go-blog/internal/models"
	"github.com/guilherme-torres/go-blog/internal/repositories"
	"github.com/guilherme-torres/go-blog/internal/services"
	_ "github.com/mattn/go-sqlite3"
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
	var email, name, password string
	flag.StringVar(&email, "email", "user@example.com", "user's email")
	flag.Parse()
	fmt.Print("Name: ")
	fmt.Scanln(&name)
	fmt.Print("Password: ")
	fmt.Scanln(&password)
	user := &models.CreateUserDTO{
		Name:     name,
		Email:    email,
		Password: password,
		Role:     "admin",
	}
	if err := userService.CreateUser(user); err != nil {
		log.Fatalln(err.Error())
	}
}
