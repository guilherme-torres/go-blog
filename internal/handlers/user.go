package handlers

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/guilherme-torres/go-blog/internal/models"
	"github.com/guilherme-torres/go-blog/internal/services"
)

type UserHandler struct {
	userService *services.UserService
	articleService     *services.ArticleService
}

func NewUserHandler(userService *services.UserService, articleService *services.ArticleService) *UserHandler {
	return &UserHandler{userService: userService, articleService: articleService}
}

func (handler *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) error {
	user := &models.CreateUserDTO{}
	user.Email = r.FormValue("email")
	user.Name = r.FormValue("name")
	user.Password = r.FormValue("password")
	user.Role = r.FormValue("role")
	err := handler.userService.CreateUser(user)
	if err != nil {
		return err
	}
	http.Redirect(w, r, "/admin/users", http.StatusSeeOther)
	return nil
}

type UsersPage struct {
	Users []*models.ListUserDTO
}

func (handler *UserHandler) ListUsers(w http.ResponseWriter, r *http.Request) error {
	users, err := handler.userService.ListUsers()
	if err != nil {
		return err
	}
	tmpl := template.Must(template.ParseFiles("./templates/users.html"))
	tmpl.Execute(w, &UsersPage{Users: users})
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
	http.Redirect(w, r, "/admin/users", http.StatusSeeOther)
	return nil
}

type AdminPage struct {
	Articles []*models.ListArticleDTO
}

func (handler *UserHandler) Admin(w http.ResponseWriter, r *http.Request) error {
	articles, err := handler.articleService.ListArticles()
	if err != nil {
		return err
	}
	tmpl := template.Must(template.ParseFiles("./templates/admin.html"))
	tmpl.Execute(w, &AdminPage{Articles: articles})
	return nil
}
