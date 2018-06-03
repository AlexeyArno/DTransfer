package network_data_handler

import (
	"sync"

	"github.com/gorilla/websocket"
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

var DataSize = 17520
var BufferSize = DataSize + 4

var PartnerConn *websocket.Conn

var (
	machines       []*Machine
	machinesLocker = &sync.Mutex{}
)

var DownloadPacketSequence chan (Packet)
var UploadPacketSequence chan (Packet)

func RegisterTCPPort(IP string) {
	TCPPort = IP
}

func RegisterPartnerConn(conn *websocket.Conn) {
	PartnerConn = conn
}
