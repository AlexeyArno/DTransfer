package websocket_work

import (
	"fmt"
	"log"

	"golang.org/x/net/websocket"
)

// socketListener listen socket for new message
// connType: 0 - info , 1 - data connection

func listenDataChannel(ws *websocket.Conn, clientOrServer uint8) {
	msg := make([]byte, 1024)
	for {
		err := websocket.Message.Receive(ws, &msg)
		if err != nil {
			log.Println("Data channel 2, type ", clientOrServer, " error: ", err)
			break
		} else {
			fmt.Printf("Received: %s.\n", msg)
		}
	}
	leaving <- ws
}

func messageGetData() {

}

func messageGetInfo() {

}
