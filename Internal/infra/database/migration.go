package database

import (
	"errors"
	"log"

	"github.com/haramurti/valeth-urls-shortener/Internal/app/user/entity"

	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(&entity.User{})
	if err != nil {
		log.Fatalf("Failed to migrate database")
		return errors.New("Failed to migrate database")
	}
	log.Println("Database Migrated Successfully")

	return nil
}
