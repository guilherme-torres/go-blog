package models

type ArticleDB struct {
	ID          int
	Title       string
	Content     string
	AuthorID    int
	PublishedAt string
	UpdatedAt   string
}

type CreateArticleDB struct {
	Title    string
	Content  string
	AuthorID int
}

type UpdateArticleDB struct {
	Title     *string
	Content   *string
	UpdatedAt string
}

type CreateArticleDTO struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type ListArticleDTO struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	AuthorID    int    `json:"authorId"`
	PublishedAt string `json:"publishedAt"`
	UpdatedAt   string `json:"updatedAt"`
}
