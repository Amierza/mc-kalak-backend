package routes

import (
	"github.com/Amierza/mc-kalak-backend/handler"
	"github.com/Amierza/mc-kalak-backend/jwt"
	"github.com/Amierza/mc-kalak-backend/middleware"
	"github.com/gin-gonic/gin"
)

func User(route *gin.Engine, userHandler handler.IUserHandler, jwtService jwt.IJWT) {
	routes := route.Group("/api/v1/users").Use(middleware.Authentication(jwtService))
	{
		routes.GET("/profile", userHandler.GetProfile)
		routes.PATCH("/profile", userHandler.Update)
	}
}
