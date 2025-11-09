package ws

import (
	"sync"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type ConnectionManager struct {
	mux         sync.Mutex
	connections map[uuid.UUID]map[uuid.UUID]websocket.Conn
}

func NewConnectionManager() *ConnectionManager {
	return &ConnectionManager{
		connections: make(map[uuid.UUID]map[uuid.UUID]websocket.Conn),
	}
}
