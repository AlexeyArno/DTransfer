package webSocketWork

import (
	"fmt"

	"golang.org/x/net/websocket"
)

// socketListener listen socket for new message
// connType: 0 - info , 1 - data connection
func socketListener(ws *websocket.Conn, connType uint8) {
	if connType == 0 {
		listenInfoChannel(ws)
	} else {
		listenDataChannel(ws)
	}
	leaving <- ws
}

func listenDataChannel(ws *websocket.Conn) {
	msg := make([]byte, 1024)
	for {
		if n, err := ws.Read(msg); err != nil {
			break
		} else {
			fmt.Printf("Received: %s.\n", msg[:n])
		}
	}
}

func listenInfoChannel(ws *websocket.Conn) {
	var msg string
	for {
		if err := websocket.Message.Receive(ws, &msg); err != nil {
			break
		} else {

		}
	}
}

func messageGetData() {

}

func messageGetInfo() {

}
