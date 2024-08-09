package main

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"

	"panel/middleware"
	"panel/utils/router"
	"panel/utils/database"
	"panel/utils/comm"
)

func init() {
	comm.Init()
    database.Init()
	
}

const (
	host     = "localhost"
	port     = 8050
)

func Init() {
    database.Init()
}

// @title Fiber API
// @version 1.0.0
// @host localhost:8050
// @BasePath /api/v1
func main() {
	app := fiber.New()
    
    middleware.Init(app)
	router.Init(app)

	Init()
	log.Fatal(app.Listen(fmt.Sprintf("%v:%v", host, port)))
}