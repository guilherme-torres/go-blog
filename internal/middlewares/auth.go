package middlewares

import (
	"context"
	"net/http"

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
			http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
			return nil
		}
		ctx := context.WithValue(r.Context(), "user_id", user.ID)
		ctx = context.WithValue(ctx, "sid", sid)
		return handler(w, r.WithContext(ctx))
	}
}
