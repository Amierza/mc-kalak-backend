package main

import (
	"log"
	"os"

	"github.com/Amierza/mc-kalak-backend/cmd"
	"github.com/Amierza/mc-kalak-backend/config/database"
	"github.com/Amierza/mc-kalak-backend/handler"
	"github.com/Amierza/mc-kalak-backend/jwt"
	"github.com/Amierza/mc-kalak-backend/middleware"
	"github.com/Amierza/mc-kalak-backend/repository"
	"github.com/Amierza/mc-kalak-backend/routes"
	"github.com/Amierza/mc-kalak-backend/service"
	"github.com/gin-gonic/gin"
)

func main() {
	db := database.SetUpPostgreSQLConnection()
	defer database.ClosePostgreSQLConnection(db)

	if len(os.Args) > 1 {
		cmd.Command(db)
		return
	}

	var (
		// jwt
		jwt = jwt.NewJWT()

		// Resource
		// User
		userRepo    = repository.NewUserRepository(db)
		userService = service.NewUserService(userRepo, jwt)
		userHandler = handler.NewUserHandler(userService)

		// Authentication
		authService = service.NewAuthService(userRepo, jwt)
		authHandler = handler.NewAuthHandler(authService)

		// Upload
		uploadService = service.NewUploadService()
		uploadHandler = handler.NewUploadHandler(uploadService)

		// Vote
		voteRepo = repository.NewVoteRepository(db)

		// Claim
		claimRepo    = repository.NewClaimRepository(db)
		claimService = service.NewClaimService(claimRepo, userRepo, voteRepo, jwt)
		claimHandler = handler.NewClaimHandler(claimService)
	)

	server := gin.Default()
	server.Use(middleware.CORSMiddleware())

	routes.User(server, userHandler, jwt)
	routes.Auth(server, authHandler, jwt)
	routes.Upload(server, uploadHandler, jwt)
	routes.Claim(server, claimHandler, jwt)

	server.Static("/uploads", "./uploads")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	var serve string
	if os.Getenv("APP_ENV") == "localhost" {
		serve = "127.0.0.1:" + port
	} else {
		serve = ":" + port
	}

	if err := server.Run(serve); err != nil {
		log.Fatalf("error running server: %v", err)
	}
}
