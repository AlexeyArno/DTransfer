package network_data_handler

import (
	"sync"

	"github.com/AlexeyArno/golang-files-transfer/src/network_client"
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

var lastPacketTimeLocker = &sync.Mutex{}
var last10PacketTime uint32 = 1
var last10PacketSize uint32 = 1
var currentPathLocker = &sync.Mutex{}
var currentPath string

// var byteTransfered uint64

var DataSize = 17520
var BufferSize = DataSize + 4

var PartnerConn *websocket.Conn

var (
	machines       []*Machine
	machinesLocker = &sync.Mutex{}
)

var DownloadPacketSequence chan (Packet) = make(chan Packet, 1)
var UploadPacketSequence chan (Packet)

func RegisterTCPPort(port string) {
	TCPPort = port
	network_client.RegTCP(port)
}

func RegisterPartnerConn(conn *websocket.Conn) {
	PartnerConn = conn
}
