package handlers

import (
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/guilherme-torres/go-blog/internal/models"
	"github.com/guilherme-torres/go-blog/internal/services"
)

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (handler *AuthHandler) Login(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "POST" {
		loginData := &models.LoginDTO{}
		err := json.NewDecoder(r.Body).Decode(loginData)
		if err != nil {
			return err
		}
		defer r.Body.Close()
		sid, err := handler.authService.Login(r.Context(), loginData)
		if err != nil {
			return err
		}
		cookie := &http.Cookie{
			Name:     "sid",
			Value:    sid,
			Path:     "/",
			MaxAge:   3600,
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteLaxMode,
		}
		http.SetCookie(w, cookie)
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
		// response := map[string]string{"sid": sid}
		// json.NewEncoder(w).Encode(response)
		return nil
	}
	tmpl := template.Must(template.ParseFiles("./assets/templates/login.html"))
	tmpl.Execute(w, nil)
	return nil
}
