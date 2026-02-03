package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/guilherme-torres/go-blog/internal/models"
	"github.com/guilherme-torres/go-blog/internal/repositories"
	"github.com/guilherme-torres/go-blog/internal/services"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/term"
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
	var email, name string
	flag.StringVar(&email, "email", "user@example.com", "user's email")
	flag.Parse()
	fmt.Print("Name: ")
	fmt.Scanln(&name)
	fmt.Print("Password: ")
	password, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Printf("\nErro ao ler senha: %v\n", err)
		os.Exit(1)
	}
	fmt.Println()
	fmt.Print("Confirm password: ")
	passwordConfirm, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Printf("\nErro ao ler senha: %v\n", err)
		os.Exit(1)
	}
	fmt.Println()
	passwordStr := string(password)
	passwordConfirmStr := string(passwordConfirm)
	if passwordConfirmStr != passwordStr{
		fmt.Println("As senhas não coincidem!")
		os.Exit(1)
	}
	user := &models.CreateUserDTO{
		Name:     name,
		Email:    email,
		Password: passwordStr,
		Role:     "admin",
	}
	if err := userService.CreateUser(user); err != nil {
		fmt.Println("Erro ao criar usuário: " + err.Error())
		os.Exit(1)
	}
}
