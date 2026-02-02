package app_errors

import "net/http"

var (
	ArticleNotFound = &AppError{ErrorCode: "article_not_found", Message: "Artigo não encontrado", StatusCode: http.StatusNotFound}
)