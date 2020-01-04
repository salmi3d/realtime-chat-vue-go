package ws

import (
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

// Client ...
type Client struct {
	ID         string
	Connection *websocket.Conn
	Pool       *Pool
}

// Message ...
type Message struct {
	ClientID string
	Body     string `json:"body"`
}

func (client *Client) Read() {
	defer func() {
		client.Pool.unregisterUser(client)
	}()

	for {
		_, p, err := client.Connection.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		message := Message{ClientID: client.ID, Body: string(p)}
		client.Pool.Broadcast <- message
		log.Printf("New message: %+v\n", message)
	}
}

// NewClient ...
func NewClient(pool *Pool, w http.ResponseWriter, r *http.Request) (*Client, error) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	client := &Client{
		ID:         uuid.New().String(),
		Connection: conn,
		Pool:       pool,
	}

	pool.registerUser(client)

	return client, nil
}
