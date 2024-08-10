package middleware

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/recover"
)

func Recover() func(c fiber.Ctx) error {
	return recover.New(recover.Config{
		Next:             nil,
		EnableStackTrace: false,
		// StackTraceHandler: defaultStackTraceHandler,
	})

	// Default
	// return recover.New(recover.Config{
	// 	Next:              nil,
	// 	EnableStackTrace:  false,
	// 	StackTraceHandler: defaultStackTraceHandler,
	// })
}
