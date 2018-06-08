package network_scan

import (
	"log"

	"github.com/AlexeyArno/golang-files-transfer/src/utility"
)

var messages = make(chan (string), 4)

//SetupSanner - Setup and start network scanner
func SetupSanner() {
	myAddress, err := utility.MyIP()
	if err != nil {
		log.Println("SetupScanner 1: ", err)
		return
	}
	port, err := ReservUDPPort()
	if err != nil {
		log.Println("SetupScanner 2: ", err)
		return
	}
	log.Println("My Address:", myAddress)
	go Listen(myAddress, port, &messages)
	go BroadcastOn(myAddress, port, &messages)
	ScanNetwork()
}

//ScanNetwork - Send broadcast request in local network with several UDP ports
func ScanNetwork() {
	for _, j := range AvailabalePortsUDP {
		messages <- "255.255.255.255:" + j
	}
}
