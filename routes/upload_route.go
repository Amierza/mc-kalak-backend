package routes

import (
	"github.com/Amierza/mc-kalak-backend/handler"
	"github.com/Amierza/mc-kalak-backend/jwt"
	"github.com/Amierza/mc-kalak-backend/middleware"
	"github.com/gin-gonic/gin"
)

func Upload(route *gin.Engine, uploadHandler handler.IUploadHandler, jwtService jwt.IJWT) {
	routes := route.Group("/api/v1/uploads").Use(middleware.Authentication(jwtService))
	{
		routes.POST("", uploadHandler.Upload)
	}
}
