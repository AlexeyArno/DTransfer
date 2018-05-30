package webSocketWork

import (
	"log"

	"github.com/AlexeyArno/golang-files-transfer/src/utility"
	"golang.org/x/net/websocket"
)

// ConnectTo data and info channels
func ConnectTo(reciverIP string) {
	myIP, err := utility.MyIP()
	if err != nil {
		log.Println("Panic myIP webSocketWork", err)
		panic(err)
	}
	if reciverIP == myIP+":"+tcpPort {
		// log.Println("Reciever IP equal my IP")
		return
	}
	origin := "http://" + reciverIP + "/"
	urlData := "ws://" + reciverIP + "/websocket_data"

	dataConn, err := websocket.Dial(urlData, "", origin)
	if err != nil {
		// log.Println("WebSocket Client 1: ", err)
	} else {
		log.Println("Connected to ", dataConn.Config().Origin.Host)
		enteringData <- dataConn
		go listenDataChannel(dataConn, 0)
	}

}
