package service

import (
	"context"
	"fmt"

	"github.com/Amierza/mc-kalak-backend/dto"
	"github.com/Amierza/mc-kalak-backend/jwt"
	"github.com/Amierza/mc-kalak-backend/repository"
	"github.com/google/uuid"
)

type (
	IUserService interface {
		GetProfile(ctx context.Context) (*dto.UserResponse, error)
		Update(ctx context.Context, req *dto.UpdateProfileRequest) (*dto.UserResponse, error)
	}

	userService struct {
		userRepo repository.IUserRepository
		jwt      jwt.IJWT
	}
)

func NewUserService(userRepo repository.IUserRepository, jwt jwt.IJWT) *userService {
	return &userService{
		userRepo: userRepo,
		jwt:      jwt,
	}
}

func (us *userService) GetProfile(ctx context.Context) (*dto.UserResponse, error) {
	token := ctx.Value("Authorization").(string)
	userIDString, err := us.jwt.GetUserIDByToken(token)
	if err != nil {
		return &dto.UserResponse{}, fmt.Errorf("Failed to get user by token: %v\n", err)
	}
	userID, err := uuid.Parse(userIDString)
	if err != nil {
		return &dto.UserResponse{}, fmt.Errorf("Failed to parse uuid into string: %v\n", err)
	}

	data, found, err := us.userRepo.GetDetailByID(ctx, nil, &userID)
	if err != nil {
		return &dto.UserResponse{}, fmt.Errorf("Failed to get user by id: %v\n", err)
	}
	if !found {
		return &dto.UserResponse{}, fmt.Errorf("Failed user not found: %v\n", err)
	}

	user := &dto.UserResponse{
		ID:        data.ID,
		Username:  data.Username,
		Password:  data.Password,
		AvatarURL: data.AvatarURL,
		IsActive:  data.IsActive,
		TimestampTemplate: dto.TimestampTemplate{
			CreatedAt: data.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: data.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	}

	return user, nil
}

func (us *userService) Update(ctx context.Context, req *dto.UpdateProfileRequest) (*dto.UserResponse, error) {
	token := ctx.Value("Authorization").(string)
	userIDString, err := us.jwt.GetUserIDByToken(token)
	if err != nil {
		return &dto.UserResponse{}, fmt.Errorf("Failed to get user ID by token: %v\n", err)
	}
	userID, err := uuid.Parse(userIDString)
	if err != nil {
		return &dto.UserResponse{}, fmt.Errorf("Failed parse id from string to uuid: %v\n", err)
	}

	user, found, err := us.userRepo.GetDetailByID(ctx, nil, &userID)
	if err != nil {
		return &dto.UserResponse{}, fmt.Errorf("Failed to get user by id: %v\n", err)
	}
	if !found {
		return &dto.UserResponse{}, fmt.Errorf("Failed user not found: %v\n", err)
	}

	user.AvatarURL = req.AvatarURL

	if err := us.userRepo.Update(ctx, nil, user); err != nil {
		return &dto.UserResponse{}, fmt.Errorf("Failed to update user: %v\n", err)
	}

	res := &dto.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Password:  user.Password,
		AvatarURL: user.AvatarURL,
		IsActive:  user.IsActive,
		TimestampTemplate: dto.TimestampTemplate{
			CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	}

	return res, nil
}
