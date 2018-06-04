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

var DownloadPacketSequence chan (Packet)
var UploadPacketSequence chan (Packet)

func RegisterTCPPort(IP string) {
	TCPPort = IP
}

func RegisterPartnerConn(conn *websocket.Conn) {
	PartnerConn = conn
}

func GetSpeed() uint32 {
	lastPacketTimeLocker.Lock()
	if last10PacketTime == 0 {
		last10PacketTime = 1
	}
	ret := (last10PacketSize / last10PacketTime) * 1000
	lastPacketTimeLocker.Unlock()
	return ret
	// return uint64(finish)
}

func GetCurrentPath() string {
	currentPathLocker.Lock()
	ret := currentPath
	currentPathLocker.Unlock()
	return ret
}
