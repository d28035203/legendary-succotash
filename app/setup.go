package app

import (
	"fmt"
	"os"

	"github.com/d28035203/legendary-succotash/database"
	"github.com/d28035203/legendary-succotash/router"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func SetupAndRunApp() error {
	_ = godotenv.Load()
	db, err := database.Connect()
	if err != nil {
		return err
	}
	app := fiber.New()
	router.SetupRoutes(app, db)
	return app.Listen(fmt.Sprintf("%s:%s", os.Getenv("APP_HOST"), os.Getenv("APP_PORT")))
}
