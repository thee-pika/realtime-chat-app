package websocket

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

// func Reader(conn *websocket.Conn) {
// 	for {
// 		messageType, p, err := conn.ReadMessage()

// 		if err != nil {
// 			log.Println(err)
// 			return
// 		}

// 		fmt.Println("msg", string(p))

// 		if err := conn.WriteMessage(messageType, p); err != nil {
// 			log.Println(err)
// 			return
// 		}
// 	}
// }

func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	fmt.Println(r.Host)

	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return conn, nil
}

// func Writer(conn *websocket.Conn) {
// 	fmt.Println("message sending .......")

// 	messageType, r, err := conn.NextReader()

// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}

// 	w, err := conn.NextWriter(messageType)

// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}

// 	if _, err := io.Copy(w, r); err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	if err := w.Close(); err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// }
