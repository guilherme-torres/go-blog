package app_errors

import "net/http"

var (
	Unauthenticated = &AppError{ErrorCode: "unauthenticated", Message: "Usuário não autenticado", StatusCode: http.StatusUnauthorized}
)