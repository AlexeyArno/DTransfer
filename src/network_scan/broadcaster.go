package network_scan

import (
	"log"
	"net"
	"sync"

	"github.com/AlexeyArno/golang-files-transfer/src/network_data_handler"
)

// BroadcastOn listen new messages and send UDP requests
func BroadcastOn(myAddres string, port string, messages *chan string) {
	// log.Println("Started broadcast")
	var wg sync.WaitGroup
	for {
		select {
		case m := <-*messages:
			wg.Add(1)
			go sendMessage(m, myAddres, port, &wg)
			wg.Wait()
		}
	}

}

// sendMessage - send UDP requests
func sendMessage(ip string, myAddres string, port string, wg *sync.WaitGroup) {
	ServerAddr, err := net.ResolveUDPAddr("udp", ip)
	if err != nil {
		log.Println("Broadcast sendMessage 1:", err)
		return
	}

	LocalAddr, err := net.ResolveUDPAddr("udp", myAddres+":"+port)
	if err != nil {
		log.Println("Broadcast sendMessage 2:", err)
		return
	}

	Conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
	if err != nil {
		log.Println("Broadcast sendMessage 3:", err)
		return
	}

	defer Conn.Close()
	_, err = Conn.Write([]byte(myAddres + ":" + network_data_handler.TCPPort))
	if err != nil {
		log.Println("Broadcast sendMessage 4:", err)
	}
	wg.Done()
}
