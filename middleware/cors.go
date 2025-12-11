package middleware

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func CorsInit(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		Next:             nil,
		AllowOriginsFunc: nil,
		AllowOrigins:     []string{ 
			// 允许请求源列表, * 表示任意请求源
			"*",
			// "http://192.168.1.100:3000"
		},
		AllowMethods: []string{
			// 允许请求方法
			fiber.MethodGet,
			fiber.MethodPost,
			fiber.MethodHead,
			fiber.MethodPut,
			fiber.MethodDelete,
			fiber.MethodPatch,
		},
		AllowHeaders:        []string{
			// 允许请求头列表, * 表示任意请求头
			"*",
		},
		AllowCredentials:    false,
		ExposeHeaders:       []string{},
		MaxAge:              0,
		AllowPrivateNetwork: false,
	}))
}