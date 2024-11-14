package polo

import (
	"sync"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Client struct {
	webSocket *websocket.Conn
	id        string
	mu        sync.Mutex
}

func NewPoloClient(webSocket *websocket.Conn) *Client {
	return &Client{
		webSocket: webSocket,
		id: func() string {
			v, _ := uuid.NewUUID()
			return v.String()
		}(),
	}
}
func (c *Client) ID() string {
	return c.id
}

func (c *Client) Send(event string, data interface{}) {
	message := map[string]interface{}{
		"event": event,
		"data":  data,
	}
	c.webSocket.WriteJSON(message)
}
