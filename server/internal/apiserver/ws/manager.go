package ws

import (
	"log"
	"sync"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type WSConnection struct {
	ID   uuid.UUID
	Conn *websocket.Conn
}

type ConnectionManager interface {
	AddConnection(userID, sessionID uuid.UUID, conn *websocket.Conn) uuid.UUID
	RemoveConnection(userID, sessionID, connectionID uuid.UUID)
}

type connectionManager struct {
	mux sync.RWMutex

	// user_id → session_id → []WSConnection
	connections map[uuid.UUID]map[uuid.UUID][]WSConnection
}

func NewConnectionManager() ConnectionManager {
	return &connectionManager{
		connections: make(map[uuid.UUID]map[uuid.UUID][]WSConnection),
	}
}

func (m *connectionManager) AddConnection(userID, sessionID uuid.UUID, conn *websocket.Conn) uuid.UUID {
	log.Printf("Adding connection for user %s, session %s\n", userID, sessionID)
	m.mux.Lock()
	defer m.mux.Unlock()

	// Generate connection ID
	connectionID := uuid.New()

	// Initialize maps if needed
	if _, exists := m.connections[userID]; !exists {
		m.connections[userID] = make(map[uuid.UUID][]WSConnection)
	}
	if _, exists := m.connections[userID][sessionID]; !exists {
		m.connections[userID][sessionID] = []WSConnection{}
	}

	// Append connection
	m.connections[userID][sessionID] = append(
		m.connections[userID][sessionID],
		WSConnection{ID: connectionID, Conn: conn},
	)

	return connectionID
}
func (m *connectionManager) RemoveConnection(userID, sessionID, connectionID uuid.UUID) {
	log.Printf("Removing connection %s for user %s, session %s\n", connectionID, userID, sessionID)
	m.mux.Lock()
	defer m.mux.Unlock()

	sessions, ok := m.connections[userID]
	if !ok {
		return
	}
	conns, ok := sessions[sessionID]
	if !ok {
		return
	}

	// Filter out the specific connection
	newList := make([]WSConnection, 0, len(conns))
	for _, c := range conns {
		if c.ID != connectionID {
			newList = append(newList, c)
		}
	}

	if len(newList) == 0 {
		delete(sessions, sessionID)
	} else {
		sessions[sessionID] = newList
	}

	if len(sessions) == 0 {
		delete(m.connections, userID)
	}
}
func (m *connectionManager) GetUserConnections(userID uuid.UUID) []WSConnection {
	m.mux.RLock()
	defer m.mux.RUnlock()

	var list []WSConnection

	sessions, ok := m.connections[userID]
	if !ok {
		return list
	}

	for _, conns := range sessions {
		list = append(list, conns...)
	}

	return list
}
func (m *connectionManager) GetSessionConnections(userID, sessionID uuid.UUID) []WSConnection {
	m.mux.RLock()
	defer m.mux.RUnlock()

	sessions, ok := m.connections[userID]
	if !ok {
		return nil
	}

	return sessions[sessionID]
}
func (m *connectionManager) BroadcastToUser(userID uuid.UUID, message any) {
	m.mux.RLock()
	connections := m.GetUserConnections(userID)
	m.mux.RUnlock()

	for _, c := range connections {
		c.Conn.WriteJSON(message)
	}
}
func (m *connectionManager) BroadcastToSession(userID, sessionID uuid.UUID, message any) {
	m.mux.RLock()
	connections := m.GetSessionConnections(userID, sessionID)
	m.mux.RUnlock()

	for _, c := range connections {
		c.Conn.WriteJSON(message)
	}
}
func (m *connectionManager) BroadcastToAll(message any) {
	m.mux.RLock()
	defer m.mux.RUnlock()

	for _, sessions := range m.connections {
		for _, conns := range sessions {
			for _, c := range conns {
				c.Conn.WriteJSON(message)
			}
		}
	}
}
