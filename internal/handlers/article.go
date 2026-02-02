package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/guilherme-torres/go-blog/internal/models"
	"github.com/guilherme-torres/go-blog/internal/services"
)

type ArticleHandler struct {
	articleService *services.ArticleService
}

func NewArticleHandler(articleService *services.ArticleService) *ArticleHandler {
	return &ArticleHandler{articleService: articleService}
}

func (handler *ArticleHandler) CreateArticle(w http.ResponseWriter, r *http.Request) error {
	userID := r.Context().Value("user_id").(int)
	article := &models.CreateArticleDTO{}
	// err := json.NewDecoder(r.Body).Decode(article)
	// if err != nil {
	// 	return err
	// }
	// defer r.Body.Close()
	article.Title = r.FormValue("title")
	article.Content = r.FormValue("content")
	err := handler.articleService.CreateArticle(article, userID)
	if err != nil {
		return err
	}
	// w.WriteHeader(http.StatusCreated)
	http.Redirect(w, r, "/admin", http.StatusSeeOther)
	return nil
}

func (handler *ArticleHandler) GetArticle(w http.ResponseWriter, r *http.Request) error {
	articleIDParam := r.PathValue("id")
	articleID, err := strconv.Atoi(articleIDParam)
	if err != nil {
		return err
	}
	article, err := handler.articleService.GetArticle(articleID)
	if err != nil {
		return err
	}
	if err := json.NewEncoder(w).Encode(article); err != nil {
		return err
	}
	return nil
}

func (handler *ArticleHandler) ListArticles(w http.ResponseWriter, r *http.Request) error {
	articles, err := handler.articleService.ListArticles()
	if err != nil {
		return err
	}
	if err := json.NewEncoder(w).Encode(articles); err != nil {
		return err
	}
	return nil
}

func (handler *ArticleHandler) DeleteArticle(w http.ResponseWriter, r *http.Request) error {
	articleIDParam := r.PathValue("id")
	articleID, err := strconv.Atoi(articleIDParam)
	if err != nil {
		return err
	}
	if err := handler.articleService.DeleteArticle(articleID); err != nil {
		return err
	}
	return nil
}
