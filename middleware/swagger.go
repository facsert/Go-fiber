package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"

	// "fibert/lib/comm"
)

func SwaggerInit(app *fiber.App) {
	// app.Use(swagger.New(swagger.Config{
	// 	BasePath: "/",
	// 	FilePath: comm.AbsPath("docs", "swagger.json"),
	// 	Path: "docs",
	// }))

	app.Get("/*", swagger.HandlerDefault)
	// app.Get("/swag/*", swagger.New(swagger.Config{
	// 	URL: "/swag/swagger.json",
	// }))
} 