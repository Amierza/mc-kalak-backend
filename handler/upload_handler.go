package handler

import (
	"net/http"

	"github.com/Amierza/mc-kalak-backend/dto"
	"github.com/Amierza/mc-kalak-backend/response"
	"github.com/Amierza/mc-kalak-backend/service"
	"github.com/gin-gonic/gin"
)

type (
	IUploadHandler interface {
		Upload(ctx *gin.Context)
	}

	uploadHandler struct {
		uploadService service.IUploadService
	}
)

func NewUploadHandler(uploadService service.IUploadService) *uploadHandler {
	return &uploadHandler{
		uploadService: uploadService,
	}
}

func (uh *uploadHandler) Upload(ctx *gin.Context) {
	file, err := ctx.FormFile("screenshot")
	if err != nil {
		res := response.BuildResponseFailed(
			dto.MESSAGE_FAILED_NO_FILES_UPLOADED,
			dto.ErrNoFilesUploaded.Error(),
			nil,
		)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	uploadedURL, err := uh.uploadService.Upload(ctx, file)
	if err != nil {
		res := response.BuildResponseFailed(
			dto.MESSAGE_FAILED_UPLOAD_FILES,
			err.Error(),
			nil,
		)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := response.BuildResponseSuccess(
		dto.MESSAGE_SUCCESS_UPLOAD_FILE,
		uploadedURL,
	)
	ctx.JSON(http.StatusOK, res)
}
