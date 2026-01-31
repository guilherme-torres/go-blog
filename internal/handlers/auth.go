package handlers

import (
	// "encoding/json"
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
		// err := json.NewDecoder(r.Body).Decode(loginData)
		// if err != nil {
		// 	return err
		// }
		// defer r.Body.Close()
		email := r.FormValue("email")
		password := r.FormValue("password")
		loginData.Email = email
		loginData.Password = password
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
		return nil
	}
	tmpl := template.Must(template.ParseFiles("./assets/templates/login.html"))
	tmpl.Execute(w, nil)
	return nil
}

func (handler *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) error {
	sid := r.Context().Value("sid").(string)
	err := handler.authService.Logout(r.Context(), sid)
	if err != nil {
		return err
	}
	cookie := &http.Cookie{
		Name:     "sid",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
	return nil
}
