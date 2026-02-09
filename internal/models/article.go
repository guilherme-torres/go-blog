package models

type ArticleDB struct {
	ID          int
	Title       string
	Content     string
	AuthorID    int
	PublishedAt string
	UpdatedAt   string
}

type ArticleWithAuthor struct {
	ID          int
	Title       string
	Content     string
	AuthorName  string
	AuthorEmail string
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
	AuthorName  string `json:"authorName"`
	AuthorEmail string `json:"authorEmail"`
	PublishedAt string `json:"publishedAt"`
	UpdatedAt   string `json:"updatedAt"`
}

type UpdateArticleDTO struct {
	Title     *string `json:"title,omitempty"`
	Content   *string `json:"content,omitempty"`
}