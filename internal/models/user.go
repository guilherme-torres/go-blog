package models

type UserDB struct {
	ID           int
	Name         string
	Email        string
	PasswordHash string
	Role         string
	CreatedAt    string
	UpdatedAt    string
}

type CreateUserDB struct {
	Name         string
	Email        string
	PasswordHash string
}

type CreateUserDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ListUserDTO struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
