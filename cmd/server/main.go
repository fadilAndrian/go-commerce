package main

import (
	"context"
	"log"
	"os"

	"github.com/fadilAndrian/go-auth/internal/handler"
	"github.com/fadilAndrian/go-auth/internal/middleware"
	"github.com/fadilAndrian/go-auth/internal/user"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	ctx := context.Background()

	db, err := pgxpool.New(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Failed to connect db, err: ", err)
	}

	r := gin.Default()

	userRepo := user.NewUserRepo(db)
	userService := user.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	authMiddleware := middleware.NewAuth()

	r.POST("/register", userHandler.Register)
	r.POST("/login", userHandler.Login)

	auth := r.Group("/auth")
	auth.Use(authMiddleware)

	auth.GET("/me", userHandler.Me)

	log.Fatal(r.Run(":8080"))
}
