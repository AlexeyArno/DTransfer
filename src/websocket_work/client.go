package websocket_work

import (
	"log"

	"github.com/AlexeyArno/golang-files-transfer/src/network_data_handler"

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
	if reciverIP == myIP+":"+network_data_handler.TCPPort {
		// log.Println("Reciever IP equal my IP")
		return
	}
	origin := "http://" + reciverIP + "/"
	urlData := "ws://" + reciverIP + "/websocket_data"

	connConfig, err := websocket.NewConfig(urlData, origin)
	if err != nil {
		log.Println("Connect to 1 : ", err)
		return
	}
	connConfig.Header.Add("Partner-Ip", myIP+":"+network_data_handler.TCPPort)

	dataConn, err := websocket.DialConfig(connConfig)
	if err != nil {
		// log.Println("WebSocket Client 1: ", err)
	} else {
		// dataConn.Config().Header.Add("Requester-IP", myIP+":"+tcpPort)
		log.Println("Connected to ", dataConn.Config().Origin.Host)
		enteringData <- dataConn
		go listenDataChannel(dataConn, 0)
	}

}
