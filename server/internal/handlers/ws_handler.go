package handlers

import (
	"log"

	"github.com/gorilla/websocket"
	"github.com/ivang5/doodle-guessr/server/internal/websockets"
	"github.com/labstack/echo/v4"
)

var upgrader = websocket.Upgrader{}

func Connect(c echo.Context) error {
	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		log.Println("Error (Connect) when upgrading connection")
		log.Printf("   |_ %v\n", err.Error())
		return err
	}

	client := websockets.NewClient(conn)
	client.Run()

	return nil
}
