package hub

import (
    "encoding/json"
    "fmt"
    "sync"

    "backendgo/models"
    "github.com/gorilla/websocket"
)

type Client struct {
    ID       string
    Conn     *websocket.Conn
    Send     chan []byte
    Partner  string // คู่แชท
}

type Hub struct {
    Clients    map[string]*Client
    mu         sync.RWMutex
    Register   chan *Client
    Unregister chan *Client
    Messages   chan *models.Message
}

func NewHub() *Hub {
    return &Hub{
        Clients:    make(map[string]*Client),
        Register:   make(chan *Client),
        Unregister: make(chan *Client),
        Messages:   make(chan *models.Message),
    }
}

func (h *Hub) Run() {
    for {
        select {
        case client := <-h.Register:
            h.mu.Lock()
            h.Clients[client.ID] = client
            h.mu.Unlock()
            fmt.Printf("✅ User %s joined\n", client.ID)

        case client := <-h.Unregister:
            h.mu.Lock()
            delete(h.Clients, client.ID)
            close(client.Send)
            h.mu.Unlock()
            fmt.Printf("❌ User %s disconnected\n", client.ID)

        case msg := <-h.Messages:
            // ส่งเฉพาะคู่แชทเท่านั้น
            h.mu.RLock()
            sender, ok1 := h.Clients[msg.From]
            receiver, ok2 := h.Clients[msg.To]
            h.mu.RUnlock()

            data, _ := json.Marshal(msg)
            if ok1 {
                sender.Send <- data // echo กลับให้ sender เห็น
            }
            if ok2 {
                receiver.Send <- data // ส่งให้คู่สนทนา
            }
        }
    }
}
