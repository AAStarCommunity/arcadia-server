package controllers

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

const ROUTER_HEALTHZ = "/ws/.internal/healthz"

func Healthz(ctx *gin.Context) {
	if ws, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil); err != nil {
		ctx.JSON(200, gin.H{
			"code": 500,
			"msg":  "failed to upgrade to websocket",
		})
	} else {
		defer func() {
			ws.Close()
		}()

		c := context.Background()

		go func() {
			<-ctx.Done()
			c.Done()
		}()

		for {
			mt, message, err := ws.ReadMessage()
			if err != nil {
				fmt.Println("read:", err)
				break
			}

			switch mt {
			case websocket.TextMessage:
				if string(message) == "ping" {
					c.Done()
					go func() {
						ws.WriteMessage(websocket.TextMessage, []byte("pong"))
					}()
				}
			}
		}
	}
}
