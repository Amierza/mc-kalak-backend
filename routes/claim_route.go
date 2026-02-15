package routes

import (
	"github.com/Amierza/mc-kalak-backend/handler"
	"github.com/Amierza/mc-kalak-backend/jwt"
	"github.com/Amierza/mc-kalak-backend/middleware"
	"github.com/gin-gonic/gin"
)

func Claim(route *gin.Engine, claimHandler handler.IClaimHandler, jwtService jwt.IJWT) {
	routes := route.Group("/api/v1/claims").Use(middleware.Authentication(jwtService))
	{
		routes.POST("", claimHandler.Create)
		routes.GET("", claimHandler.GetAll)
		routes.GET("/:id", claimHandler.GetDetailByID)
		routes.PUT("/:id", claimHandler.Update)
		routes.DELETE("/:id", claimHandler.DeleteByID)

		// Vote
		routes.POST("/:id/vote", claimHandler.Vote)
		routes.GET("/:id/vote", claimHandler.GetAllVotesByClaimID)
	}
}
