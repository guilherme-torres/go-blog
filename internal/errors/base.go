package app_errors

type ApiError struct {
	StatusCode int
	ErrorCode  string
	Message    string
}
