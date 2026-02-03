package handlers

import (
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"

	"github.com/guilherme-torres/go-blog/internal/models"
	"github.com/guilherme-torres/go-blog/internal/services"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (handler *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) error {
	user := &models.CreateUserDTO{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	err = handler.userService.CreateUser(user)
	if err != nil {
		return err
	}
	w.WriteHeader(http.StatusCreated)
	return nil
}

func (handler *UserHandler) ListUsers(w http.ResponseWriter, r *http.Request) error {
	users, err := handler.userService.ListUsers()
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(users); err != nil {
		return err
	}
	return nil
}

func (handler *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) error {
	userIDParam := r.PathValue("id")
	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		return err
	}
	if err := handler.userService.DeleteUser(userID); err != nil {
		return err
	}
	return nil
}

func (handler *UserHandler) Admin(w http.ResponseWriter, r *http.Request) error {
	tmpl := template.Must(template.ParseFiles("./assets/templates/admin.html"))
	tmpl.Execute(w, nil)
	return nil
}
