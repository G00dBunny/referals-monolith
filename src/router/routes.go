package routes

import (
	"referals/src/controllers"

	"github.com/gofiber/fiber/v3"
)

func Setup(app *fiber.App){
	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Server is working! 🚀")
	})
	api := app.Group("api")
	admin := api.Group("admin")

	admin.Post("/register", controllers.Register)
	admin.Post("/login", controllers.Login)
}