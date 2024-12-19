package database

import (
	"fmt"
	"github.com/joramuns/shop/internal/models"
	"gorm.io/gorm"
	"log"
)

func RunUserMigration(db *gorm.DB) error {
	err := db.AutoMigrate(&models.User{})
	if err != nil {
		return fmt.Errorf("AutoMigrate err: %v", err)
	}
	log.Println("Database migration completed")
	return nil
}
