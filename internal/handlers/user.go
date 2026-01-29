package handlers

import (
	"encoding/json"
	"fmt"
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

func (handler *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	user := &models.CreateUserDTO{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		http.Error(w, "", 500)
		return
	}
	err = handler.userService.CreateUser(user)
	if err != nil {
		http.Error(w, "", 400)
		return
	}
	w.WriteHeader(201)
	fmt.Fprintf(w, "")
}

func (handler *UserHandler) ListUsers(w http.ResponseWriter, r *http.Request) {

}

func (handler *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {

}

func (handler *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {

}
