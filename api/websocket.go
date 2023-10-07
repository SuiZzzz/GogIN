package api

import (
	"GoGin/service"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandleClient() gin.HandlerFunc {
	return func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Println("websocket err:", err)
			return
		}
		client := service.NewClient()
		go func() {
			for {
				resp := client.HandleMessage(conn, c)
				bytes, _ := json.Marshal(resp)
				err = conn.WriteMessage(websocket.TextMessage, bytes)
				if err != nil {
					break
				}
			}
		}()
	}
}
