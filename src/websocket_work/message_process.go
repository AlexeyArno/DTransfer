package websocket_work

import (
	"fmt"
	"log"
	"strings"

	"golang.org/x/net/websocket"
)

// socketListener listen socket for new message
// connType: 0 - info , 1 - data connection

func listenDataChannel(ws *websocket.Conn, clientOrServer uint8) {
	log.Println("Goroutin 2 spawn")
	msg := make([]byte, 1024)
	count := 0
	for {
		err := websocket.Message.Receive(ws, &msg)
		if err != nil {
			log.Println("Data channel 2, type ", clientOrServer, " error: ", err)
			break
		} else {
			if count == 0 {
				m := machineIsHere(ws)
				if m != nil {
					if len(strings.Split(m.IP, ":")) == 1 {
						m.IP += ":" + string(msg)
					}
				} else {
					log.Println("Fuck")
				}
				count++
			}

			fmt.Printf("Received: %s.\n", msg)
		}
	}
	leaving <- ws
}

func messageGetData() {

}

func messageGetInfo() {

}
