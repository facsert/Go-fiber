package socket

import (
    // "encoding/json"
    "fmt"
    
	"github.com/gofiber/fiber/v2"
    "github.com/gofiber/contrib/socketio"
    "github.com/gofiber/fiber/v2/log"
)

var (
	clients = make(map[string]string, 10)
)


func Init(api fiber.Router) {
    api.Get("/socket.io/:id", socketio.New(func(kws *socketio.Websocket) {

        // Retrieve the user id from endpoint
        userId := kws.Params("id")

        // Add the connection to the list of the connected clients
        // The UUID is generated randomly and is the key that allow
        // socketio to manage Emit/EmitTo/Broadcast
        clients[userId] = kws.UUID

        // Every websocket connection has an optional session key => value storage
        kws.SetAttribute("user_id", userId)

        //Broadcast to all the connected users the newcomer
        kws.Broadcast([]byte(fmt.Sprintf("New user connected: %s and UUID: %s", userId, kws.UUID)), true, socketio.TextMessage)
        //Write welcome message
        kws.Emit([]byte(fmt.Sprintf("Hello user: %s with UUID: %s", userId, kws.UUID)), socketio.TextMessage)
    }))

    // Multiple event handling supported
    socketio.On(socketio.EventConnect, func(ep *socketio.EventPayload) {
        log.Infof(fmt.Sprintf("Connection event 1 - User: %s", ep.Kws.GetStringAttribute("user_id")))
    })

    // On disconnect event
    socketio.On(socketio.EventDisconnect, func(ep *socketio.EventPayload) {
        // Remove the user from the local clients
        delete(clients, ep.Kws.GetStringAttribute("user_id"))
        log.Infof(fmt.Sprintf("Disconnection event - User: %s", ep.Kws.GetStringAttribute("user_id")))
    })

    // On close event
    // This event is called when the server disconnects the user actively with .Close() method
    socketio.On(socketio.EventClose, func(ep *socketio.EventPayload) {
        // Remove the user from the local clients
        delete(clients, ep.Kws.GetStringAttribute("user_id"))
        log.Infof(fmt.Sprintf("Close event - User: %s", ep.Kws.GetStringAttribute("user_id")))
    })

    // On error event
    socketio.On(socketio.EventError, func(ep *socketio.EventPayload) {
        log.Infof(fmt.Sprintf("Error event - User: %s", ep.Kws.GetStringAttribute("user_id")))
    })

    socketio.On("command", func(ep *socketio.EventPayload) {
        log.Infof(fmt.Sprintf("customer event - User: %s", ep.Kws.GetStringAttribute("command")))
    })
}
