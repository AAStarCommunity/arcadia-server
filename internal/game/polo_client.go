package game

import (
	"sync"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type PoloClient struct {
	webSocket *websocket.Conn
	id        string
	mu        sync.Mutex
}

func NewPoloClient(webSocket *websocket.Conn) *PoloClient {
	return &PoloClient{
		webSocket: webSocket,
		id: func() string {
			v, _ := uuid.NewUUID()
			return v.String()
		}(),
	}
}
func (c *PoloClient) ID() string {
	return c.id
}
