package routes

import (
	"github.com/Amierza/go-boiler-plate/handler"
	"github.com/Amierza/go-boiler-plate/jwt"
	"github.com/Amierza/go-boiler-plate/middleware"
	"github.com/gin-gonic/gin"
)

func Upload(route *gin.Engine, uploadHandler handler.IUploadHandler, jwtService jwt.IJWT) {
	routes := route.Group("/api/v1/uploads").Use(middleware.Authentication(jwtService))
	{
		routes.POST("", uploadHandler.Upload)
	}
}
