package routes

import (
	"github.com/Amierza/mc-kalak-backend/handler"
	"github.com/Amierza/mc-kalak-backend/jwt"
	"github.com/gin-gonic/gin"
)

func User(route *gin.Engine, userHandler handler.IUserHandler, jwtService jwt.IJWT) {
	routes := route.Group("/api/v1/users")
	{
		routes.Use()
	}
}
