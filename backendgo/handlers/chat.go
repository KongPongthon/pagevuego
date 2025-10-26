package handlers

import (
    "encoding/json"
    "log"
    "net/http"

    "backendgo/hub"
    "backendgo/models"
    "github.com/gin-gonic/gin"
    "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool { return true },
}

func ServeWs(h *hub.Hub, c *gin.Context) {
    userID := c.Query("user")
    partner := c.Query("partner")

    conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
    if err != nil {
        log.Println("Upgrade error:", err)
        return
    }

    client := &hub.Client{
        ID:      userID,
        Conn:    conn,
        Send:    make(chan []byte),
        Partner: partner,
    }

    h.Register <- client

    // อ่านจาก client
    go func() {
        defer func() {
            h.Unregister <- client
            conn.Close()
        }()
        for {
            _, msg, err := conn.ReadMessage()
            if err != nil {
                log.Println("Read error:", err)
                break
            }

            var message models.Message
            if err := json.Unmarshal(msg, &message); err == nil {
                h.Messages <- &message
            }
        }
    }()

    // ส่งข้อความออกไปให้ client
    go func() {
        for msg := range client.Send {
            if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
                log.Println("Write error:", err)
                break
            }
        }
    }()
}
