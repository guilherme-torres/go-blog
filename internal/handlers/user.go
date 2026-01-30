package handlers

import (
	"encoding/json"
	"net/http"

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
