package ws

import (
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
	pool.trackUser(client, "New User Joined...")
}

func (pool *Pool) unregisterUser(client *Client) {
	if err := client.Connection.Close(); err != nil {
		log.Println(err)
	}
	delete(pool.Clients, client.ID)
	pool.trackUser(client, "User Disconnected...")
}

func (pool *Pool) trackUser(client *Client, str string) {
	log.Println("Connected clients: ", len(pool.Clients))

	// var dat map[string]interface{}
	// if err := json.Unmarshal([]byte(str), &dat); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println("---", dat["type"])

	msg := Message{
		ClientID: client.ID,
		Body:     str,
	}
	pool.Broadcast <- msg
}

// Discover ...
func (pool *Pool) Discover() {
	for {
		select {
		case message := <-pool.Broadcast:
			log.Println("Broadcasting:", message)
			for _, client := range pool.Clients {
				if message.ClientID == client.ID {
					continue
				}
				if err := client.Connection.WriteJSON(message); err != nil {
					log.Println(err)
					return
				}
			}
		}
	}
}
