package ws

import (
	"fmt"
	"log"
)

// Pool ...
type Pool struct {
	Clients   map[string]*Client
	Broadcast chan Message
}

// NewPool ...
func NewPool() *Pool {
	return &Pool{
		Clients:   make(map[string]*Client),
		Broadcast: make(chan Message),
	}
}

func (pool *Pool) registerUser(client *Client) {
	pool.Clients[client.ID] = client
	pool.trackUser(client, "+")
}

func (pool *Pool) unregisterUser(client *Client) {
	if err := client.Connection.Close(); err != nil {
		log.Println(err)
	}
	pool.trackUser(client, "-")
}

func (pool *Pool) trackUser(client *Client, str string) {
	log.Printf("Connected clients: %d\n", len(pool.Clients))
	var msgBody, msgType string
	if str == "+" {
		msgBody = fmt.Sprintf("%s joined...", client.Name)
	} else if str == "-" {
		msgBody = fmt.Sprintf("%s disconnected...", client.Name)
		msgType = "unregister"
	}
	message := Message{ClientID: client.ID, Type: msgType, Body: msgBody}
	pool.Broadcast <- message
}

// Discover ...
func (pool *Pool) Discover() {
	for {
		select {
		case message := <-pool.Broadcast:
			log.Println("Broadcasting:", message)
			author := pool.Clients[message.ClientID].Name
			if message.Type == "unregister" {
				delete(pool.Clients, message.ClientID)
			}
			for _, client := range pool.Clients {
				if message.ClientID == client.ID {
					continue
				}
				response := map[string]string{"author": author, "text": message.Body}
				err := client.Connection.WriteJSON(response)
				if err != nil {
					log.Println(err)
					continue
				}
			}
		}
	}
}
