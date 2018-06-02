package websocket_work

import "golang.org/x/net/websocket"

type machine struct {
	dataConn *websocket.Conn
	IP       string
}
