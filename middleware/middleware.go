package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func Init(app *fiber.App) {
    // app.Use(PanicHandle)
	app.Use(Logger())
	app.Use(Recover())
	app.Use(CorsInit())
	// app.Use(SocketInit())
}