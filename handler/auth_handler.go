package handler

import (
	"net/http"

	"github.com/Amierza/mc-kalak-backend/dto"
	"github.com/Amierza/mc-kalak-backend/response"
	"github.com/Amierza/mc-kalak-backend/service"
	"github.com/gin-gonic/gin"
)

type (
	IAuthHandler interface {
		Login(ctx *gin.Context)
	}

	authHandler struct {
		authService service.IAuthService
	}
)

func NewAuthHandler(authService service.IAuthService) *authHandler {
	return &authHandler{
		authService: authService,
	}
}

func (ah *authHandler) Login(ctx *gin.Context) {
	var payload dto.LoginRequest
	if err := ctx.ShouldBind(&payload); err != nil {
		res := response.BuildResponseFailed(dto.MESSAGE_INVALID_REQUEST_PAYLOAD, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := ah.authService.Login(ctx, payload)
	if err != nil {
		res := response.BuildResponseFailed(dto.FAILED_LOGIN, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := response.BuildResponseSuccess(dto.SUCCESS_LOGIN, result)
	ctx.JSON(http.StatusOK, res)
}
