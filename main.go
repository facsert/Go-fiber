package main

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"

	"panel/middleware"
	"panel/utils/comm"
	"panel/utils/database"
	"panel/utils/router"
)


const (
	host = "localhost"
	port = 8050
)

func Init(app *fiber.App) {
	comm.Init()
	middleware.Init(app)
	router.Init(app)

	database.Init()
}

// @title Fiber API
// @version 1.0.0
// @host localhost:8050
// @BasePath /api/v1
func main() {
	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName: "Test App v1.0.1",
	})
	Init(app)
	log.Fatal(app.Listen(fmt.Sprintf("%v:%v", host, port), fiber.ListenConfig{
		EnablePrefork: true,
		DisableStartupMessage: true,
	}))
}
