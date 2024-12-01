package main

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type WebSocketServer struct {
	clients    map[*websocket.Conn]bool // List of connected clients
	broadcast  chan []byte              // Channel for broadcast messages
	upgrader   websocket.Upgrader       // WebSocket upgrader
	clientsMux sync.Mutex               // Mutex for synchronizing access to clients
}

func NewWebSocketServer() *WebSocketServer {
	return &WebSocketServer{
		clients:   make(map[*websocket.Conn]bool),
		broadcast: make(chan []byte),
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool { return true },
		},
	}
}

func (s *WebSocketServer) handleConnections(w http.ResponseWriter, r *http.Request) {
	// Upgrade HTTP connection to WebSocket
	conn, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Failed to upgrade to WebSocket:", err)
		return
	}
	defer conn.Close()

	// Add the new client to the list
	s.clientsMux.Lock()
	s.clients[conn] = true
	s.clientsMux.Unlock()
	log.Println("New client connected")

	// Receive messages from the client
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Client disconnected:", err)
			s.clientsMux.Lock()
			delete(s.clients, conn)
			s.clientsMux.Unlock()
			break
		}
		log.Printf("Message received: %s\n", message)

		// Send the message to the broadcast channel
		s.broadcast <- message
	}
}

func (s *WebSocketServer) handleBroadcast() {
	for {
		// Receive messages from the broadcast channel
		message := <-s.broadcast

		// Send the message to all clients
		s.clientsMux.Lock()
		for client := range s.clients {
			err := client.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				log.Println("Failed to send message:", err)
				client.Close()
				delete(s.clients, client)
			}
		}
		s.clientsMux.Unlock()
	}
}

func main() {
	server := NewWebSocketServer()

	// WebSocket endpoint
	http.HandleFunc("/ws", server.handleConnections)

	// Run a goroutine to handle broadcast
	go server.handleBroadcast()

	// Start the HTTP server
	addr := "localhost:8080"
	log.Printf("Server running on %s...\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal("Failed to start the server:", err)
	}
}
