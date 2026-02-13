package handler

import (
	"github.com/Amierza/mc-kalak-backend/service"
	"go.uber.org/zap"
)

type (
	IUserHandler interface {
	}

	userHandler struct {
		userService service.IUserService
		logger      *zap.Logger
	}
)

func NewUserHandler(userService service.IUserService, logger *zap.Logger) *userHandler {
	return &userHandler{
		userService: userService,
		logger:      logger,
	}
}
