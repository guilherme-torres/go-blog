package repositories

import (
	"database/sql"

	"github.com/guilherme-torres/go-blog/internal/models"
)

type ArticleRepository struct {
	db *sql.DB
}

func NewArticleRepo(db *sql.DB) *ArticleRepository {
	return &ArticleRepository{db: db}
}

func (repo *ArticleRepository) Create(article *models.CreateArticleDB) (int64, error) {
	result, err := repo.db.Exec(`
		INSERT OR IGNORE INTO "articles" ("title", "content", "author_id") VALUES (?, ?, ?)
	`, article.Title, article.Content, article.AuthorID)
	if err != nil {
		return 0, err
	}
	rowAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowAffected, nil
}

func (repo *ArticleRepository) List() ([]*models.ArticleDB, error) {
	rows, err := repo.db.Query(`SELECT "id", "title", "content", "author_id", "published_at", "updated_at" FROM "articles"`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []*models.ArticleDB
	for rows.Next() {
		article := &models.ArticleDB{}
		if err := rows.Scan(
			&article.ID,
			&article.Title,
			&article.Content,
			&article.AuthorID,
			&article.PublishedAt,
			&article.UpdatedAt,
		); err != nil {
			return nil, err
		}
		users = append(users, article)
	}
	return users, nil
}

func (repo *ArticleRepository) Get(id int) (*models.ArticleDB, error) {
	row := repo.db.QueryRow(`
		SELECT "id", "title", "content", "author_id", "published_at", "updated_at"
		FROM "users" WHERE "id" = ?`, id,
	)
	article := &models.ArticleDB{}
	err := row.Scan(
		&article.ID,
		&article.Title,
		&article.Content,
		&article.AuthorID,
		&article.PublishedAt,
		&article.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return article, nil
}

func (repo *ArticleRepository) Update(id int, article *models.ArticleDB) {

}

func (repo *ArticleRepository) Delete(id int) (int64, error) {
	result, err := repo.db.Exec(`
		DELETE FROM "articles" WHERE "id" = ?
	`, id)
	if err != nil {
		return 0, err
	}
	rowAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowAffected, nil
}
