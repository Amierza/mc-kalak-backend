package routes

import (
	"github.com/Amierza/mc-kalak-backend/handler"
	"github.com/Amierza/mc-kalak-backend/jwt"
	"github.com/gin-gonic/gin"
)

func Auth(route *gin.Engine, authHandler handler.IAuthHandler, jwtService jwt.IJWT) {
	routes := route.Group("/api/v1/auth")
	{
		routes.POST("/login", authHandler.Login)
	}
}
