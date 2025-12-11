package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v3"

	"devops/pkgs/comm"
	"devops/pkgs/db"
	"devops/api/user"
	"devops/middleware"
)

var (
	HOST = "0.0.0.0"
	PORT = 3100
	APP  *fiber.App
)

func main() {
	APP = NewApp()
	defer Listen()

	comm.Init()
	db.Init()

	InitMiddle()
	InitRouters()
}

func NewApp() *fiber.App {
	return fiber.New(fiber.Config{
		ServerHeader:  "Develop Server",
		AppName:       "Develop V1.0.0.0",
		CaseSensitive: true,
		StrictRouting: false,
		BodyLimit:     4 * 1024 * 1024,
	})
}

func Listen() {
	log.Fatal(APP.Listen(
		fmt.Sprintf("%s:%d", HOST, PORT),
		fiber.ListenConfig{
			EnablePrefork:         true,
			DisableStartupMessage: false,
			EnablePrintRoutes:     true,
		},
	))
}

func InitMiddle() {
	middleware.CorsInit(APP)
	middleware.LoggerInit(APP)
	middleware.RecoverInit(APP)
}

func InitRouters() {
	router := APP.Group("api/v1")
	user.Init(router)
}

