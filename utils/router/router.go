package router

import (
	"github.com/gofiber/fiber/v3"

	"panel/api/v1/node"
	// "fibert/api/v1/scan"
)

func Init(app *fiber.App) {
	api := app.Group("/api/v1")
    
    node.Init(api)
    // article.Init(api)
	// scan.Init(api)
}
