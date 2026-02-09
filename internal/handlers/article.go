package handlers

import (
	"html/template"
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
	article.Title = r.FormValue("title")
	article.Content = r.FormValue("content")
	err := handler.articleService.CreateArticle(article, userID)
	if err != nil {
		return err
	}
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
	tmpl := template.Must(template.ParseFiles(
		"./templates/articles-base.html",
		"./templates/article-detail.html",
	))
	tmpl.ExecuteTemplate(w, "base", article)
	return nil
}

type ArticlesPage struct {
	Articles []*models.ListArticleDTO
}

func (handler *ArticleHandler) ListArticles(w http.ResponseWriter, r *http.Request) error {
	articles, err := handler.articleService.ListArticles()
	if err != nil {
		return err
	}
	tmpl := template.Must(template.ParseFiles(
		"./templates/articles-base.html",
		"./templates/articles.html",
	))
	tmpl.ExecuteTemplate(w, "base", &ArticlesPage{Articles: articles})
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
	http.Redirect(w, r, "/admin", http.StatusSeeOther)
	return nil
}

func (handler *ArticleHandler) UpdateArticle(w http.ResponseWriter, r *http.Request) error {
	articleIDParam := r.PathValue("id")
	articleID, err := strconv.Atoi(articleIDParam)
	if err != nil {
		return err
	}
	article := &models.UpdateArticleDTO{}
	titleStr := r.FormValue("title")
	contentStr := r.FormValue("content")
	article.Title = &titleStr
	article.Content = &contentStr
	if err := handler.articleService.UpdateArticle(articleID, article); err != nil {
		return err
	}
	return nil
}
