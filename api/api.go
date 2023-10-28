package api

import (
	"os"

	"github.com/ayowilfred95/config"
	"github.com/ayowilfred95/database"
	"github.com/ayowilfred95/model"
	"github.com/ayowilfred95/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Api() error {

	// init env
	err := config.LoadEnv()
	if err != nil {
		return err
	}

	// init database
	database.ConnectDB()
	database.DB.AutoMigrate(&model.User{})

	// Create an app using Fiber framework
	app := fiber.New()


	// Initialize cors middleware default config to allow all application access 
	//app.Use(cors.New())

	// Initialize cors middleware with specific allowed origins
	// Initialize cors middleware to allow all origins
	app.Use(cors.New(cors.Config{
    AllowOrigins: "*",
    AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// add basic middleware
	app.Use(logger.New())
	app.Use(recover.New())

	// add routes
	router.SetupRoutes(app)

	// start the server
	var port string

	if port = os.Getenv("PORT"); port == "" {
		port = "8080"
	}
	app.Listen(":" + port)
	return nil

}
