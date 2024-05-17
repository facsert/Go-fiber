package router

import (
	"github.com/gofiber/fiber/v2"

	"fibert/api/v1/article"
	"fibert/api/v1/scan"
)

func Init(app *fiber.App) {
	api := app.Group("/api/v1")

    article.Init(api)
	scan.Init(api)
}

