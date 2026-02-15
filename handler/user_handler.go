package handler

import (
	"fmt"
	"net/http"

	"github.com/Amierza/mc-kalak-backend/dto"
	"github.com/Amierza/mc-kalak-backend/response"
	"github.com/Amierza/mc-kalak-backend/service"
	"github.com/gin-gonic/gin"
)

type (
	IUserHandler interface {
		GetProfile(ctx *gin.Context)
		Update(ctx *gin.Context)
	}

	userHandler struct {
		userService service.IUserService
	}
)

func NewUserHandler(userService service.IUserService) *userHandler {
	return &userHandler{
		userService: userService,
	}
}

func (uh *userHandler) GetProfile(ctx *gin.Context) {
	result, err := uh.userService.GetProfile(ctx)
	if err != nil {
		res := response.BuildResponseFailed(fmt.Sprintf("%s user", dto.FAILED_GET_PROFILE), err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := response.BuildResponseSuccess(fmt.Sprintf("%s user", dto.SUCCESS_GET_PROFILE), result)
	ctx.JSON(http.StatusOK, res)
}

func (uh *userHandler) Update(ctx *gin.Context) {
	payload := &dto.UpdateProfileRequest{}
	if err := ctx.ShouldBind(&payload); err != nil {
		res := response.BuildResponseFailed(dto.MESSAGE_INVALID_REQUEST_PAYLOAD, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := uh.userService.Update(ctx, payload)
	if err != nil {
		res := response.BuildResponseFailed(fmt.Sprintf("%s user", dto.FAILED_UPDATE), err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := response.BuildResponseSuccess(fmt.Sprintf("%s user", dto.SUCCESS_UPDATE), result)
	ctx.JSON(http.StatusOK, res)
}
