package services

import (
	"context"
	"encoding/hex"
	"strconv"
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
	expTime := 40 * time.Minute
	if err := service.redisClient.Set(ctx, "session:" + sidHashString, user.ID, expTime); err != nil {
		return "", err
	}
	return sidString, nil
}

func (service *AuthService) DeleteSession(ctx context.Context, sid string) error {
	sidHashBytes, err := utils.Sha256Hash(sid)
	if err != nil {
		return err
	}
	sidHashString := hex.EncodeToString(sidHashBytes)
	if err := service.redisClient.Del(ctx, "session:" + sidHashString); err != nil {
		return err
	}
	return nil
}

func (service *AuthService) VerifySession(ctx context.Context, sid string) (*models.UserDB, error) {
	sidHashBytes, err := utils.Sha256Hash(sid)
	if err != nil {
		return nil, err
	}
	sidHashString := hex.EncodeToString(sidHashBytes)
	value, err := service.redisClient.Get(ctx, "session:" + sidHashString)
	if err != nil {
		return nil, err
	}
	if value == "" {
		return nil, app_errors.Unauthenticated
	}
	userID, err := strconv.Atoi(value)
    if err != nil {
        return nil, err
    }
	user, err := service.userRepo.Get(userID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, app_errors.Unauthenticated
	}
	// extende o tempo de expiração da sessão após uma interação do usuário
	expTime := 40 * time.Minute
	if err := service.redisClient.Set(ctx, "session:" + sidHashString, user.ID, expTime); err != nil {
		return nil, err
	}
	return user, nil
}
