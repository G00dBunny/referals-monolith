package main

import (
	"log"
	"referals/src/database"
	routes "referals/src/router"

	"github.com/gofiber/fiber/v3"
)

func main() {

	database.Connect()

	database.AutoMigrate()
    // Initialize a new Fiber app
    app := fiber.New()

   	routes.Setup(app)

    log.Fatal(app.Listen(":8000"))
}