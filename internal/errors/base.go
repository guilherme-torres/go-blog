package app_errors

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type AppError struct {
	StatusCode int    `json:"status_code"`
	ErrorCode  string `json:"error_code"`
	Message    string `json:"message"`
}

func (e *AppError) Error() string {
	return fmt.Sprintf("[%s] %s", e.ErrorCode, e.Message)
}

type Handler func(w http.ResponseWriter, r *http.Request) error

func HandleErrors(handler Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := handler(w, r); err != nil {
			var appError *AppError
			if errors.As(err, &appError) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(appError.StatusCode)
				json.NewEncoder(w).Encode(appError)
				return
			} else {
				internalError := map[string]any{
					"statusCode": http.StatusInternalServerError,
					"message":    "internal server error",
				}
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(internalError)
				return
			}
		}
	}
}
