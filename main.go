package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/yusong-offx/myChat/chat"
)

func main() {
	app := fiber.New()

	// get websocket request
	app.Use("/ws/", chat.RequestWebsocket)

	// make websocket
	app.Get("/ws/:user_id", websocket.New(chat.ChatApp))

	log.Fatal(app.Listen(":8080"))
}
