package network_data_handler

import (
	"sync"

	"github.com/AlexeyArno/golang-files-transfer/src/network_client"
	"github.com/gorilla/websocket"
)

var TCPPort string

var (
	UploadPath   string
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

//DownloadPacketSequence , UploadPacketSequence - channels for send and get packets between machines
var DownloadPacketSequence = make(chan Packet, 1)
var UploadPacketSequence chan (Packet)

// UploadCallback , DownloadCallback Functions for call after Upload or Download done
var UploadCallback func()
var DownloadCallback func()

// RegisterTCPPort set TCP port for our app
func RegisterTCPPort(port string) {
	TCPPort = port
	network_client.RegTCP(port)
}

// RegisterPartnerConn set reciever or sender IP
// depending on load type
func RegisterPartnerConn(conn *websocket.Conn) {
	PartnerConn = conn
}
