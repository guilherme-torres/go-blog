package app_errors

import "fmt"

type UserError struct {
	ApiError
}

func (e *UserError) Error() string {
	return fmt.Sprintf("[%s] %s", e.ErrorCode, e.Message)
}

var (
	GenericUserError = &UserError{ApiError{ErrorCode: "user_error", Message: "Erro", StatusCode: 500}}
	UserAlreadyExists = &UserError{ApiError{ErrorCode: "user_already_exists", Message: "Este usuário já existe", StatusCode: 400}}
	InvalidCredentials = &UserError{ApiError{ErrorCode: "invalid_credentials", Message: "Credenciais inválidas", StatusCode: 400}}
	UserNotFound = &UserError{ApiError{ErrorCode: "user_not_found", Message: "Usuário não encontrado", StatusCode: 404}}
)