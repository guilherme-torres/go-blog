package services

import (
	app_errors "github.com/guilherme-torres/go-blog/internal/errors"
	"github.com/guilherme-torres/go-blog/internal/models"
	"github.com/guilherme-torres/go-blog/internal/repositories"
	"github.com/guilherme-torres/go-blog/internal/utils"
)

type ArticleService struct {
	articleRepo *repositories.ArticleRepository
}

func NewArticleService(articleRepo *repositories.ArticleRepository) *ArticleService {
	return &ArticleService{articleRepo: articleRepo}
}

func (service *ArticleService) CreateArticle(article *models.CreateArticleDTO, userID int) error {
	newArticle := &models.CreateArticleDB{
		Title:    article.Title,
		Content:  article.Content,
		AuthorID: userID,
	}
	_, err := service.articleRepo.Create(newArticle)
	if err != nil {
		return err
	}
	return nil
}

func (service *ArticleService) GetArticle(id int) (*models.ListArticleDTO, error) {
	article, err := service.articleRepo.Get(id)
	if err != nil {
		return nil, err
	}
	if article == nil {
		return nil, app_errors.ArticleNotFound
	}
	return &models.ListArticleDTO{
		ID:          article.ID,
		Title:       article.Title,
		Content:     article.Content,
		AuthorID:    article.AuthorID,
		PublishedAt: article.PublishedAt,
		UpdatedAt:   article.UpdatedAt,
	}, nil
}

func (service *ArticleService) ListArticles() ([]*models.ListArticleDTO, error) {
	articles, err := service.articleRepo.List()
	if err != nil {
		return nil, err
	}
	articlesResponse := utils.Map(articles, func(article *models.ArticleDB) *models.ListArticleDTO {
		return &models.ListArticleDTO{
			ID:          article.ID,
			Title:       article.Title,
			Content:     article.Content,
			AuthorID:    article.AuthorID,
			PublishedAt: article.PublishedAt,
			UpdatedAt:   article.UpdatedAt,
		}
	})
	return articlesResponse, nil
}

func (service *ArticleService) UpdateArticle(id int) (*models.ListArticleDTO, error) {
	return nil, nil
}

func (service *ArticleService) DeleteArticle(id int) error {
	rowsAffected, err := service.articleRepo.Delete(id)
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return app_errors.ArticleNotFound
	}
	return nil
}
