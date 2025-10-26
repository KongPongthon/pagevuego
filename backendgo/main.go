package main

import (
    "backendgo/handlers"
    "backendgo/hub"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    h := hub.NewHub()

    go h.Run() // รัน hub background

    r.GET("/ws", func(c *gin.Context) {
        handlers.ServeWs(h, c)
    })

    r.Run(":8080")
}