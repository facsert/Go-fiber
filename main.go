package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/swagger"

	"fibert/lib/database"
	"fibert/lib/router"
	"fibert/lib/comm"
	"fibert/middleware"
	_ "fibert/docs"
)

func init() {
	comm.Init()
    database.Init()
}

const (
	host     = "localhost"
	port     = 8050
)

// @title Fiber API
// @version 1.0.0
// @host localhost:8050
// @BasePath /api/v1
func main() {
	app := fiber.New()
    
    middleware.Init(app)
	router.Init(app)
	app.Get("/*", swagger.HandlerDefault)
	
	log.Fatal(app.Listen(fmt.Sprintf("%v:%v", host, port)))
}

// app := fiber.New(fiber.Config{
// 	Prefork:       true,
// 	CaseSensitive: true,
// 	StrictRouting: true,
// 	ServerHeader:  "Fiber",
// 	AppName: "Test App v1.0.1",
// 	BodyLimit: 300 * 1024 * 1024, // 上传文件限制 300M
// })