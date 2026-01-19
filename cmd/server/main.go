package main

import (
	"context"
	"log"
	"os"

	"github.com/fadilAndrian/go-commerce/internal/handler"
	"github.com/fadilAndrian/go-commerce/internal/middleware"
	"github.com/fadilAndrian/go-commerce/internal/product"
	"github.com/fadilAndrian/go-commerce/internal/user"
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

	productRepo := product.NewProductRepo(db)
	productService := product.NewProductService(productRepo)
	productHandler := handler.NewProductHandler(productService)

	authMiddleware := middleware.NewAuth()

	r.POST("/register", userHandler.Register)
	r.POST("/login", userHandler.Login)

	auth := r.Group("/auth")
	auth.Use(authMiddleware)
	auth.GET("/me", userHandler.Me)

	product := r.Group("/products")
	product.Use(authMiddleware)
	product.GET("/", productHandler.List)
	product.POST("/", productHandler.Create)
	product.GET("/:id", productHandler.Show)
	product.PUT("/:id", productHandler.Update)
	product.DELETE("/:id", productHandler.Delete)

	log.Fatal(r.Run(":8080"))
}
