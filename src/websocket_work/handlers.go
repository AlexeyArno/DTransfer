package websocket_work

import (
	"log"

	"github.com/AlexeyArno/golang-files-transfer/src/network_data_handler"
	"github.com/AlexeyArno/golang-files-transfer/src/utility"
	"golang.org/x/net/websocket"
)

var (
	enteringData = make(chan *websocket.Conn)
	leaving      = make(chan *websocket.Conn)
)

func init() {
	go WebsocketService()
}

func DataHandler(ws *websocket.Conn) {

	log.Println("Client origin: ", ws.Request().RemoteAddr, "conn to data-channel")
	enteringData <- ws
	listenDataChannel(ws, 1)
}

// WebsocketService process all new and old websocket connections
func WebsocketService() {
	myIP, err := utility.MyIP()
	if err != nil {
		return
	}
	for {
		select {

		case c := <-enteringData:
			if c == nil {
				log.Println("Data channel faile")
				return
			}
			if network_data_handler.MachineIsHere(c) {
				log.Println("Websocket Work 1:", "machine already exist")
				c.Close()
				continue
			}
			ip := c.Config().Origin.Host
			if ip == myIP+":"+network_data_handler.TCPPort {
				ip = c.Request().Header.Get("Partner-Ip")
			}
			if ip == "" {
				log.Println("Websocket Work 2:", "ip is empty")
				c.Close()
				continue
			}
			network_data_handler.AddMachine(ip, c)
		case c := <-leaving:
			if network_data_handler.MachineIsHere(c) {
				network_data_handler.DeleteMachine(c)
			}
		}
	}
}
