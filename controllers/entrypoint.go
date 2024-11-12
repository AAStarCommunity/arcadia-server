package controllers

import (
	"arcadia/internal/events"
	"context"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type genericEvent struct {
	Event string          `json:"event"`
	Data  json.RawMessage `json:"data"`
}

func Entrypoint(ctx *gin.Context) {
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
				} else {
					if eventName, err := parserMessage(message); err == nil {
						switch eventName.Event {
						case events.JoinEventName:
							// if joinData, err := parseJoinEventData(eventName); err == nil {
							// go services.JoinGame(joinData, ws)
							// }
						}
					}
				}
			}
		}
	}
}

func parserMessage(message []byte) (*genericEvent, error) {
	evt := genericEvent{}
	err := json.Unmarshal(message, &evt)

	return &evt, err
}

func parseJoinEventData(ge *genericEvent) (*events.JoinEventData, error) {
	evt := events.JoinEventData{
		Name: ge.Event,
	}
	err := json.Unmarshal(ge.Data, &evt)

	return &evt, err
}
