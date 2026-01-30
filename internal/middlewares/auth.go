package middlewares

import (
	"fmt"
	"net/http"

	app_errors "github.com/guilherme-torres/go-blog/internal/errors"
)

func AuthMiddleware(handler app_errors.Handler) app_errors.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		cookie, err := r.Cookie("sid")
		if err != nil {
			if err == http.ErrNoCookie {
				http.Redirect(w, r, "/auth/login", http.StatusSeeOther)
				return app_errors.Unauthenticated
			}
			return err
		}
		sid := cookie.Value
		fmt.Println("sid:", sid)
		return handler(w, r)
	}
}
