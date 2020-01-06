package ws

import (
	"errors"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

// Client ...
type Client struct {
	ID         string
	Name       string
	Connection *websocket.Conn
	Pool       *Pool
}

// Message ...
type Message struct {
	ClientID string
	Type     string
	Body     string
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

	name, ok := r.URL.Query()["name"]

	if !ok || len(name[0]) < 1 {
		return nil, errors.New("Url Param 'name' is missing")
	}

	clientName := name[0]

	client := &Client{
		ID:         uuid.New().String(),
		Name:       clientName,
		Connection: conn,
		Pool:       pool,
	}

	pool.registerUser(client)

	return client, nil
}
