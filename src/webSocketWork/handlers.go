package webSocketWork

import (
	"log"

	gui "github.com/AlexeyArno/golang-files-transfer/src/gui"
	"golang.org/x/net/websocket"
)

func DataHandler(ws *websocket.Conn) {
	gui.Log("Client conn to data-channel: " + ws.RemoteAddr().String())
	log.Println("Client conn to data-channel: ", ws.RemoteAddr().String())
}

func InfoHandler(ws *websocket.Conn) {
	gui.Log("Client conn to info-channel: " + ws.RemoteAddr().String())
	log.Println("Client conn to info-channel: ", ws.RemoteAddr().String())
}
