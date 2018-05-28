package network

import (
	"log"
)

func StartScanNetworks() {
	log.Println("Start Listeneing")
	go BroadcastToNetwork()
	go Listen()
}
