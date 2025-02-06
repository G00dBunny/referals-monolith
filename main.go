package main

import (
	"log"
	"referals/src/database"
	routes "referals/src/router"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func main() {


	database.Connect()

	database.AutoMigrate()
    // Initialize a new Fiber app
    app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins: []string{"http://localhost:8000"},
	}))
	
   	routes.Setup(app)

    log.Fatal(app.Listen(":8000"))
}