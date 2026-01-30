package app_errors

import "net/http"

var (
	GenericUserError   = &AppError{ErrorCode: "user_error", Message: "Erro", StatusCode: http.StatusInternalServerError}
	UserAlreadyExists  = &AppError{ErrorCode: "user_already_exists", Message: "Este usuário já existe", StatusCode: http.StatusBadRequest}
	InvalidCredentials = &AppError{ErrorCode: "invalid_credentials", Message: "Credenciais inválidas", StatusCode: http.StatusBadRequest}
	UserNotFound       = &AppError{ErrorCode: "user_not_found", Message: "Usuário não encontrado", StatusCode: http.StatusNotFound}
)