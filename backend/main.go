package main

import (
	"log"
	"net/http"

	"github.com/salmi3d/realtime-chat-vue-go/backend/pkg/ws"
)

func initEndpoints() {
	pool := ws.NewPool()
	go pool.Discover()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		client, err := ws.NewClient(pool, w, r)
		if err != nil {
			return
		}
		client.Read()
	})
}

func main() {
	log.Println("Starting chat daemon...")
	initEndpoints()
	log.Fatal(http.ListenAndServe(":4444", nil))
}
