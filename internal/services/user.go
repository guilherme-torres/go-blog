package services

import (
	"github.com/guilherme-torres/go-blog/internal/errors"
	"github.com/guilherme-torres/go-blog/internal/models"
	"github.com/guilherme-torres/go-blog/internal/repositories"
	"github.com/guilherme-torres/go-blog/internal/utils"
)

type UserService struct {
	userRepo *repositories.UserRepository
}

func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (service *UserService) CreateUser(user *models.CreateUserDTO) error {
	passwordHash, err := utils.HashPassword(user.Password)
	if err != nil {
		return app_errors.InvalidCredentials
	}
	newUser := &models.CreateUserDB{
		Name:         user.Name,
		Email:        user.Email,
		PasswordHash: passwordHash,
	}
	rowsAffected, err := service.userRepo.Create(newUser)
	if err != nil {
		return app_errors.GenericUserError
	}
	if rowsAffected == 0 {
		return app_errors.UserAlreadyExists
	}
	return nil
}

func (service *UserService) ListUsers() {

}

func (service *UserService) GetUser(id int) (*models.ListUserDTO, error) {
	user, err := service.userRepo.Get(id)
	if err != nil {
		return nil, app_errors.UserNotFound
	}
	return &models.ListUserDTO{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
	}, nil
}

func (service *UserService) DeleteUser(id int) error {
	rowsAffected, err := service.userRepo.Delete(id)
	if err != nil {
		return app_errors.GenericUserError
	}
	if rowsAffected == 0 {
		return app_errors.UserNotFound
	}
	return nil
}
