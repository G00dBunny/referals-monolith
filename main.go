package main

import (
	"log"
	"referals/src/database"
<<<<<<< HEAD
	routes "referals/src/router"
=======
>>>>>>> 7df7561 (added mysql db)

	"github.com/gofiber/fiber/v3"
)

func main() {

	database.Connect()

<<<<<<< HEAD
	database.AutoMigrate()
    // Initialize a new Fiber app
    app := fiber.New()

   	routes.Setup(app)
=======
    // Initialize a new Fiber app
    app := fiber.New()

    // Define a route for the GET method on the root path '/'
    app.Get("/", func(c fiber.Ctx) error {
        // Send a string response to the client
        return c.SendString("Hello, World 👋!")
    })
>>>>>>> 7df7561 (added mysql db)

    log.Fatal(app.Listen(":8000"))
}