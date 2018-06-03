package websocket_work

import (
	"log"

	"github.com/gorilla/websocket"
)

// socketListener listen socket for new message
// connType: 0 - info , 1 - data connection

func listenDataChannel(ws *websocket.Conn, clientOrServer uint8) {
	// msg := make([]byte, 4100)
	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			log.Println("Data channel 2, type ", clientOrServer, " error: ", err)
			break
		} else {
			log.Println("Received: ", msg[:4], " ", msg[5:8])
			// ip, err := network_data_handler.GetIPByConnection(ws)
			// if err != nil {
			// 	continue
			// } else {
			// 	network_data_handler.DownloadPacketSequence <- network_data_handler.Packet{IP: ip, Data: msg}
			// }

		}
	}
	leaving <- ws
}

func messageGetData() {

}

func messageGetInfo() {

}
