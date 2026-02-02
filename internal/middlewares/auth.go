package middlewares

import (
	"context"
	"html/template"
	"net/http"
	"slices"

	app_errors "github.com/guilherme-torres/go-blog/internal/errors"
	"github.com/guilherme-torres/go-blog/internal/services"
)

func AuthMiddleware(handler app_errors.Handler, authService *services.AuthService) app_errors.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		cookie, err := r.Cookie("sid")
		if err != nil {
			if err == http.ErrNoCookie {
				http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
				return nil
			}
			return err
		}
		sid := cookie.Value
		user, err := authService.VerifySession(r.Context(), sid)
		if err != nil {
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
		// TODO: armazenar um hashmap contendo o "user_id", "sid" e outros dados relevantes
		ctx := context.WithValue(r.Context(), "user_id", user.ID)
		ctx = context.WithValue(ctx, "sid", sid)
		ctx = context.WithValue(ctx, "user_role", user.Role)
		return handler(w, r.WithContext(ctx))
	}
}

func VerifyRole(handler app_errors.Handler, roles []string, authService *services.AuthService) app_errors.Handler {
	errorTmpl := template.Must(template.ParseFiles("./assets/templates/403.html"))
	return func(w http.ResponseWriter, r *http.Request) error {
		userRole := r.Context().Value("user_role").(string)
		sid := r.Context().Value("sid").(string)
		if !slices.Contains(roles, userRole) {
			if err := authService.DeleteSession(r.Context(), sid); err != nil {
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
			errorTmpl.Execute(w, nil)
			return nil
		}
		return handler(w, r)
	}
}
