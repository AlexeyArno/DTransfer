package network_data_handler

import (
	"sync"

	"golang.org/x/net/websocket"
)

var TCPPort string

var (
	UploadPath   string = `D:\GOPATH\src\github.com\AlexeyArno\golang-files-transfer\src`
	DownloadPath string
)

var (
	RecieverIP string
	SenderIP   string
)

var PartnerConn *websocket.Conn

var (
	machines       []*Machine
	machinesLocker = &sync.Mutex{}
)

func RegisterTCPPort(IP string) {
	TCPPort = IP
}

func RegisterPartnerConn(conn *websocket.Conn) {
	PartnerConn = conn
}
