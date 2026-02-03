package services

import (
	app_errors "github.com/guilherme-torres/go-blog/internal/errors"
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
		return err
	}
	newUser := &models.CreateUserDB{
		Name:         user.Name,
		Email:        user.Email,
		PasswordHash: passwordHash,
		Role:         user.Role,
	}
	rowsAffected, err := service.userRepo.Create(newUser)
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return app_errors.UserAlreadyExists
	}
	return nil
}

func (service *UserService) ListUsers() ([]*models.ListUserDTO, error) {
	users, err := service.userRepo.List()
	if err != nil {
		return nil, err
	}
	usersResponse := utils.Map(users, func(user *models.UserDB) *models.ListUserDTO {
		return &models.ListUserDTO{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		}
	})
	return usersResponse, nil
}

func (service *UserService) GetUser(id int) (*models.ListUserDTO, error) {
	user, err := service.userRepo.Get(id)
	if err != nil {
		return nil, app_errors.UserNotFound
	}
	return &models.ListUserDTO{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (service *UserService) DeleteUser(id int) error {
	rowsAffected, err := service.userRepo.Delete(id)
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return app_errors.UserNotFound
	}
	return nil
}
