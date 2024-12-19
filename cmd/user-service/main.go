package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joramuns/shop/internal/database"
	"github.com/joramuns/shop/internal/handlers"
	"github.com/joramuns/shop/internal/kafka"
	"github.com/joramuns/shop/internal/repository"
	"github.com/joramuns/shop/internal/service"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	log.Printf("Connecting to database %s", dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := database.RunUserMigration(db); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	repo := repository.NewUserRepository(db)
	producer := kafka.NewProducer([]string{"kafka:9092"}, "user_updates")
	defer producer.Close()

	userService := service.NewUserService(repo, producer)
	userHandler := handlers.NewUserHandler(userService)

	router := gin.Default()
	router.POST("/users", userHandler.RegisterUser)
	router.GET("/list", userHandler.ListUsers)

	log.Println("Server is running on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
