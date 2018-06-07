package network_scan

import (
	"log"

	"github.com/AlexeyArno/golang-files-transfer/src/utility"
)

var messages chan (string) = make(chan string, 4)

func SetupSanner() {
	myAddress, err := utility.MyIP()
	if err != nil {
		log.Println(err)
		return
	}
	port, err := ReservUDPPort()
	if err != nil {
		log.Println("Setup: ", err)
		return
	}
	// log.Println("My Address:", myAddress)
	go Listen(myAddress, port, IPFound, &messages)
	go BroadcastTo(myAddress, port, &messages)
	ScanNetwork()

}

func ScanNetwork() {
	for _, j := range AvailabalePortsUDP {
		messages <- "255.255.255.255:" + j
	}
}

func IPFound(ip string) {
	// log.Println("Ip Found: ", ip)
}
