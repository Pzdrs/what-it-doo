package controller

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"pycrs.cz/what-it-doo/internal/apiserver/ws"
)

type SocketController struct {
	upgrader          websocket.Upgrader
	connectionManager *ws.ConnectionManager
}

func NewSocketController(upgrader websocket.Upgrader, connectionManager *ws.ConnectionManager) *SocketController {
	return &SocketController{upgrader: upgrader, connectionManager: connectionManager}
}

func (c *SocketController) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := c.upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Failed to upgrade to WebSocket", http.StatusInternalServerError)
		return
	}

	go func() {
		defer conn.Close()
		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				break
			}
			// Handle incoming WebSocket messages
			fmt.Println("Received WebSocket message:", string(msg))
		}
	}()
}
