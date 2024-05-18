package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func CorsInit() func(c *fiber.Ctx) error {
	return cors.New(cors.Config{
        AllowOrigins: "*", // 允许所有域
        AllowHeaders: "Origin, Content-Type, Accept, Authorization",
        AllowMethods: "GET, POST, HEAD, PUT, DELETE, PATCH, OPTIONS",
	})
}