package app_errors

var (
	GenericUserError = &AppError{ErrorCode: "user_error", Message: "Erro", StatusCode: 500}
	UserAlreadyExists = &AppError{ErrorCode: "user_already_exists", Message: "Este usuário já existe", StatusCode: 400}
	InvalidCredentials = &AppError{ErrorCode: "invalid_credentials", Message: "Credenciais inválidas", StatusCode: 400}
	UserNotFound = &AppError{ErrorCode: "user_not_found", Message: "Usuário não encontrado", StatusCode: 404}
)