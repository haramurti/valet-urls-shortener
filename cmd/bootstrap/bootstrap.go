package bootstrap

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/haramurti/valeth-urls-shortener/Internal/infra/database"
	"github.com/haramurti/valeth-urls-shortener/config"
	"gorm.io/gorm"
)

type App struct {
	Fiber  *fiber.App
	DB     *gorm.DB
	Config *config.Config
}

func NewApp() *App {
	cfg := config.Load()

	db, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database")
		panic(err)
	}

	//Migration
	err = database.AutoMigrate(db)
	if err != nil {
		log.Fatalf("Failed to migrate database")
		panic(err)
	}
	return &App{
		Fiber:  fiber.New(),
		DB:     db,
		Config: cfg,
	}
}
