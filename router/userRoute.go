package router

import (
	"github.com/ayowilfred95/controller"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	users := app.Group("/api")

	users.Post("/waitlist", controller.CreateUser)
	users.Post("/waitlist/user", controller.CreateUserWithEmail)
	users.Put("/waitlist/user",controller.UpdateUser)
	users.Get("/waitlist", controller.GetUsers)
}
