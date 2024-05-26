package middleware

import (
	"github.com/gofiber/fiber/v2"
    // "github.com/gofiber/contrib/socketio"
    "github.com/gofiber/contrib/websocket"
)

func SocketInit() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
        // IsWebSocketUpgrade returns true if the client
        // requested upgrade to the WebSocket protocol.
        if websocket.IsWebSocketUpgrade(c) {
            c.Locals("allowed", true)
            return c.Next()
        }
        return fiber.ErrUpgradeRequired
    }
}



