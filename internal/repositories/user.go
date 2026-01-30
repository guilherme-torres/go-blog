package repositories

import (
	"database/sql"

	"github.com/guilherme-torres/go-blog/internal/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (repo *UserRepository) Create(user *models.CreateUserDB) (int64, error) {
	result, err := repo.db.Exec(`
		INSERT OR IGNORE INTO "users" ("name", "email", "password_hash") VALUES (?, ?, ?)
	`, user.Name, user.Email, user.PasswordHash)
	if err != nil {
		return 0, err
	}
	rowAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowAffected, nil
}

func (repo *UserRepository) List() ([]*models.UserDB, error) {
	rows, err := repo.db.Query(`SELECT "id", "name", "email", "password_hash", "role", "created_at", "updated_at" FROM "users"`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []*models.UserDB
	for rows.Next() {
		user := &models.UserDB{}
		if err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.PasswordHash,
			&user.Role,
			&user.CreatedAt,
			&user.UpdatedAt,
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (repo *UserRepository) Get(id int) (*models.UserDB, error) {
	row := repo.db.QueryRow(`
		SELECT "id", "name", "email", "password_hash", "role", "created_at", "updated_at"
		FROM "users" WHERE "id" = ?`, id,
	)
	user := &models.UserDB{}
	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.PasswordHash,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) FindByEmail(email string) (*models.UserDB, error) {
	row := repo.db.QueryRow(`
		SELECT "id", "name", "email", "password_hash", "role", "created_at", "updated_at"
		FROM "users" WHERE "email" = ?`, email,
	)
	user := &models.UserDB{}
	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.PasswordHash,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) Update(id int, user *models.UserDB) {

}

func (repo *UserRepository) Delete(id int) (int64, error) {
	result, err := repo.db.Exec(`
		DELETE FROM "users" WHERE "id" = ?
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
