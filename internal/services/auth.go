package services

import (
	"context"
	"encoding/hex"
	"time"

	app_errors "github.com/guilherme-torres/go-blog/internal/errors"
	"github.com/guilherme-torres/go-blog/internal/models"
	"github.com/guilherme-torres/go-blog/internal/repositories"
	"github.com/guilherme-torres/go-blog/internal/utils"
)

type AuthService struct {
	userRepo    *repositories.UserRepository
	redisClient *utils.RedisClient
}

func NewAuthService(userRepo *repositories.UserRepository, redisClient *utils.RedisClient) *AuthService {
	return &AuthService{userRepo: userRepo, redisClient: redisClient}
}

func (service *AuthService) Login(ctx context.Context, data *models.LoginDTO) (string, error) {
	user, err := service.userRepo.FindByEmail(data.Email)
	if err != nil {
		return "", err
	}
	if user == nil || !utils.VerifyPasswordHash(data.Password, user.PasswordHash) {
		return "", app_errors.InvalidCredentials
	}
	sidBytes, err := utils.GenerateRandomBytes(32)
	if err != nil {
		return "", err
	}
	sidString := hex.EncodeToString(sidBytes)
	sidHashBytes, err := utils.Sha256Hash(sidString)
	if err != nil {
		return "", err
	}
	sidHashString := hex.EncodeToString(sidHashBytes)
	expTime := 1 * time.Minute
	if err := service.redisClient.Set(ctx, "session:" + sidHashString, user.ID, expTime); err != nil {
		return "", err
	}
	return sidString, nil
}
