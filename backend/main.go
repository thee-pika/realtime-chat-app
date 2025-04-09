package main

import (
	"fmt"
	"net/http"

	"github.com/thee-pika/realtime-chat-app/pkg/websocket"
)

func serverWsHandler(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {

	conn, err := websocket.Upgrade(w, r)

	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}

	client := &websocket.Client{
		Conn: conn,
		Pool: pool,
	}

	pool.Register <- client
	client.Read()
}

func getRequestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Server called for get Request")
}

func setUpRoutes() {
	http.HandleFunc("/", getRequestHandler)

	pool := websocket.NewPool()
	go pool.Start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serverWsHandler(pool, w, r)
	})
}

func main() {
	setUpRoutes()
	fmt.Println("server is running on port", 8000)
	http.ListenAndServe(":8000", nil)
}
