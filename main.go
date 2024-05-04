package main

import (
	"fmt"
	
	"github.com/gofiber/fiber/v2"
    // "github.com/gofiber/swagger"

	"fibert/lib/logger"
	"fibert/lib/database"
	"fibert/lib/router"
	"fibert/middleware/swagger"
    _ "fibert/docs"
)

func init() {
    logger.InitLogger()
    database.InitDatabase()
}

const (
	host     = "192.168.1.100"
	port     = 8010
)

// @title Fiber API
// @version 1.0.0
// @host 192.168.1.100:8010
// @BasePath /api/v1
func main() {
	app := fiber.New()
    
	router.InitRouter(app)
	swag.SwaggerInit(app)
	
	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("Hello, World!")
	// })

	app.Listen(fmt.Sprintf("%v:%v", host, port))
}