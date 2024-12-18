package user_service

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/joramuns/shop/internal/handlers"
	"github.com/joramuns/shop/internal/kafka"
	"github.com/joramuns/shop/internal/repository"
	"github.com/joramuns/shop/internal/service"
	"log"
)

func main() {
	dsn := "host=db user=postgres password=postgres dbname=users port=5432 sslmode=disable"
	db, err := gorm.Open(dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	repo := repository.NewUserRepository(db)
	producer := kafka.NewProducer([]string{"kafka:9092"}, "user_updates")
	defer producer.Close()

	userService := service.NewUserService(repo, producer)
	userHandler := handlers.NewUserHandler(userService)

	router := gin.Default()
	router.POST("/users", userHandler.RegisterUser)

	log.Println("Server is running on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
