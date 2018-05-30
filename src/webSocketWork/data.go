package webSocketWork

import "golang.org/x/net/websocket"

type machine struct {
	dataConn *websocket.Conn
	IP       string
}

var tcpPort string

func RegisterTCPPort(IP string) {
	tcpPort = IP
}
