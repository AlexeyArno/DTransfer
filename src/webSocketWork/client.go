package webSocketWork

import (
	"log"

	"golang.org/x/net/websocket"
)

// ConnectTo data and info channels
func ConnectTo(ip string) {
	origin := "http://" + ip + "/"
	urlData := "ws://" + ip + "/websocket_data"
	urlInfo := "ws://" + ip + "/websocket_info"

	_, err := websocket.Dial(urlData, "", origin)
	if err != nil {
		// log.Println("WebSocket Client 1: ", err)
	} else {
		log.Println("Connected to ", urlData)
	}

	_, err = websocket.Dial(urlInfo, "", origin)
	if err != nil {
		// log.Println("WebSocket Client 2: ", err)
	} else {
		log.Println("Connected to ", urlInfo)
	}

}
