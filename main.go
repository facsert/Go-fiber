package main

import (
	"github.com/gofiber/fiber/v2"

	"fibert/lib/logger"
	"fibert/lib/database"
	"fibert/lib/router"

)

func init() {
    logger.InitLogger()
    database.InitDatabase()
}

func main() {
	app := fiber.New()
    
	router.InitRouter(app)
	
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}