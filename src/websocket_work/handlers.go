package websocket_work

import (
	"log"
	"net/http"

	"github.com/AlexeyArno/golang-files-transfer/src/network_data_handler"
	"github.com/AlexeyArno/golang-files-transfer/src/utility"
	"github.com/gorilla/websocket"
)

type newConn struct {
	conn *websocket.Conn
	ip   string
}

var (
	myIP         string
	enteringData = make(chan newConn)
	leaving      = make(chan *websocket.Conn)
)

func init() {
	ip, err := utility.MyIP()
	if err != nil {
		return
	}
	myIP = ip
	go WebsocketService()
}

func DataHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Upgrade(w, r, nil, network_data_handler.BufferSize, network_data_handler.BufferSize)
	if err != nil {
		log.Println("Data Handler 1 ", err)
		return
	}
	ip := r.Host
	if len(ip) == 0 {
		return
	}
	if ip == myIP+":"+network_data_handler.TCPPort {
		ip = r.Header.Get("Partner-Ip")
	}

	// log.Println("Client origin: ", Request().RemoteAddr, "conn to data-channel")
	enteringData <- newConn{conn: conn, ip: ip}
	listenDataChannel(conn, 1)
}

// WebsocketService process all new and old websocket connections
func WebsocketService() {

	for {
		select {

		case c := <-enteringData:
			if c.conn == nil {
				log.Println("Data channel faile")
				return
			}
			if network_data_handler.MachineIsHere(c.conn) {
				log.Println("Websocket Work 1:", "machine already exist")
				c.conn.Close()
				continue
			}
			network_data_handler.AddMachine(c.ip, c.conn)
		case c := <-leaving:
			if network_data_handler.MachineIsHere(c) {
				network_data_handler.DeleteMachine(c)
			}
		}
	}
}
