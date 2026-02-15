package handler

import (
	"fmt"
	"net/http"

	"github.com/Amierza/mc-kalak-backend/dto"
	"github.com/Amierza/mc-kalak-backend/response"
	"github.com/Amierza/mc-kalak-backend/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type (
	IClaimHandler interface {
		Create(ctx *gin.Context)
		GetAll(ctx *gin.Context)
		GetDetailByID(ctx *gin.Context)
		Update(ctx *gin.Context)
		DeleteByID(ctx *gin.Context)

		Vote(ctx *gin.Context)
		GetAllVotesByClaimID(ctx *gin.Context)
	}

	claimHandler struct {
		claimService service.IClaimService
	}
)

func NewClaimHandler(claimService service.IClaimService) *claimHandler {
	return &claimHandler{
		claimService: claimService,
	}
}

func (ch *claimHandler) Create(ctx *gin.Context) {
	payload := &dto.CreateClaimRequest{}
	if err := ctx.ShouldBind(&payload); err != nil {
		res := response.BuildResponseFailed(dto.MESSAGE_INVALID_REQUEST_PAYLOAD, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := ch.claimService.Create(ctx, payload)
	if err != nil {
		res := response.BuildResponseFailed(fmt.Sprintf("%s claim", dto.FAILED_CREATE), err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := response.BuildResponseSuccess(fmt.Sprintf("%s claim", dto.SUCCESS_CREATE), result)
	ctx.JSON(http.StatusOK, res)
}

func (ch *claimHandler) GetAll(ctx *gin.Context) {
	var pagination response.PaginationRequest
	if err := ctx.ShouldBindQuery(&pagination); err != nil {
		res := response.BuildResponseFailed(dto.MESSAGE_INVALID_QUERY_PARAMS, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := ch.claimService.GetAllWithPagination(ctx, pagination)
	if err != nil {
		res := response.BuildResponseFailed(fmt.Sprintf("%s claims", dto.FAILED_GET_ALL), err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := response.Response{
		Status:   true,
		Messsage: fmt.Sprintf("%s claims", dto.SUCCESS_GET_ALL),
		Data:     result.Data,
		Meta:     result.PaginationResponse,
	}
	ctx.JSON(http.StatusOK, res)
}

func (ch *claimHandler) GetDetailByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		res := response.BuildResponseFailed(dto.MESSAGE_INVALID_QUERY_PARAMS, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := ch.claimService.GetDetailByID(ctx, &id)
	if err != nil {
		res := response.BuildResponseFailed(fmt.Sprintf("%s claim", dto.FAILED_GET_DETAIL), err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := response.BuildResponseSuccess(fmt.Sprintf("%s claim", dto.SUCCESS_GET_DETAIL), result)
	ctx.JSON(http.StatusOK, res)
}

func (ch *claimHandler) Update(ctx *gin.Context) {
	payload := &dto.UpdateClaimRequest{}
	idStr := ctx.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		res := response.BuildResponseFailed(dto.MESSAGE_INVALID_QUERY_PARAMS, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	payload.ID = id

	if err := ctx.ShouldBind(&payload); err != nil {
		res := response.BuildResponseFailed(dto.MESSAGE_INVALID_REQUEST_PAYLOAD, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := ch.claimService.Update(ctx, payload)
	if err != nil {
		res := response.BuildResponseFailed(fmt.Sprintf("%s claim", dto.FAILED_UPDATE), err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := response.BuildResponseSuccess(fmt.Sprintf("%s claim", dto.SUCCESS_UPDATE), result)
	ctx.JSON(http.StatusOK, res)
}

func (ch *claimHandler) DeleteByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		res := response.BuildResponseFailed(dto.MESSAGE_INVALID_QUERY_PARAMS, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := ch.claimService.DeleteByID(ctx, &id)
	if err != nil {
		res := response.BuildResponseFailed(fmt.Sprintf("%s claim", dto.FAILED_DELETE), err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := response.BuildResponseSuccess(fmt.Sprintf("%s claim", dto.SUCCESS_DELETE), result)
	ctx.JSON(http.StatusOK, res)
}

func (ch *claimHandler) Vote(ctx *gin.Context) {
	payload := &dto.ClaimVoteRequest{}
	idStr := ctx.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		res := response.BuildResponseFailed(dto.MESSAGE_INVALID_QUERY_PARAMS, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	payload.ID = id

	if err := ctx.ShouldBind(&payload); err != nil {
		res := response.BuildResponseFailed(dto.MESSAGE_INVALID_REQUEST_PAYLOAD, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := ch.claimService.Vote(ctx, payload)
	if err != nil {
		res := response.BuildResponseFailed("failed to vote claim", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := response.BuildResponseSuccess("success to vote claim", result)
	ctx.JSON(http.StatusOK, res)
}

func (ch *claimHandler) GetAllVotesByClaimID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		res := response.BuildResponseFailed(dto.MESSAGE_INVALID_QUERY_PARAMS, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := ch.claimService.GetAllVotesByClaimID(ctx, &id)
	if err != nil {
		res := response.BuildResponseFailed(fmt.Sprintf("%s votes", dto.FAILED_GET_ALL), err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := response.BuildResponseSuccess(fmt.Sprintf("%s votes", dto.SUCCESS_GET_ALL), result)
	ctx.JSON(http.StatusOK, res)
}
