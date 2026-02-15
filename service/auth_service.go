package service

import (
	"context"
	"fmt"

	"github.com/Amierza/mc-kalak-backend/dto"
	"github.com/Amierza/mc-kalak-backend/helper"
	"github.com/Amierza/mc-kalak-backend/jwt"
	"github.com/Amierza/mc-kalak-backend/repository"
)

type (
	IAuthService interface {
		Login(ctx context.Context, req dto.LoginRequest) (dto.LoginResponse, error)
	}

	authService struct {
		userRepo repository.IUserRepository
		jwt      jwt.IJWT
	}
)

func NewAuthService(userRepo repository.IUserRepository, jwt jwt.IJWT) *authService {
	return &authService{
		userRepo: userRepo,
		jwt:      jwt,
	}
}

func (as *authService) Login(ctx context.Context, req dto.LoginRequest) (dto.LoginResponse, error) {
	user, found, err := as.userRepo.GetByUsername(ctx, nil, &req.Username)
	if err != nil {
		return dto.LoginResponse{}, fmt.Errorf("Failed to get user by username: %v\n", err)
	}
	if !found {
		return dto.LoginResponse{}, fmt.Errorf("User not found with username: %s\n", req.Username)
	}

	checkPassword, err := helper.CheckPassword(user.Password, []byte(req.Password))
	if err != nil {
		return dto.LoginResponse{}, fmt.Errorf("Failed to check password: %v\n", err)
	}
	if !checkPassword {
		return dto.LoginResponse{}, fmt.Errorf("Incorrect password for username: %s\n", req.Username)
	}

	token, err := as.jwt.GenerateToken(user.ID.String())
	if err != nil {
		return dto.LoginResponse{}, fmt.Errorf("Failed to generate token for userID %s: %v\n", user.ID.String(), err)
	}

	return dto.LoginResponse{
		Token: token,
	}, nil
}
