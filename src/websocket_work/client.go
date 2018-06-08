package websocket_work

import (
	"log"
	"net/http"

	"github.com/AlexeyArno/golang-files-transfer/src/network_data_handler"

	"net/url"

	"github.com/AlexeyArno/golang-files-transfer/src/utility"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  network_data_handler.BufferSize,
	WriteBufferSize: network_data_handler.BufferSize,
}

// ConnectTo data and info channels
func ConnectTo(reciverIP string) {
	myIP, err := utility.MyIP()
	if err != nil {
		log.Println("Panic myIP webSocketWork", err)
		panic(err)
	}
	if reciverIP == myIP+":"+network_data_handler.TCPPort {
		return
	}

	u := url.URL{Scheme: "ws", Host: reciverIP, Path: "/websocket_data"}
	log.Println("try connect to ", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(),
		http.Header{"Partner-Ip": []string{myIP + ":" + network_data_handler.TCPPort}})
	if err != nil {
		log.Fatal("Dial Client Websocket_Work:", err)
		return
	}

	enteringData <- newConn{c, reciverIP}
	go listenDataChannel(c, 0)

}
