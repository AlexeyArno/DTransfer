package network_scan

import (
	"log"
	"net"

	websocket_work "github.com/AlexeyArno/golang-files-transfer/src/websocket_work"
)

var (
	usefullIps []string
)

// Listen UDP requests
func Listen(myAddressWPort string, port string, messages *chan string) {

	ServerAddr, err := net.ResolveUDPAddr("udp", ":"+port)
	if err != nil {
		log.Println("Listen 1:", err)
		return
	}

	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	if err != nil {
		log.Println("Listen 2:", err)
		return
	}
	defer ServerConn.Close()
	log.Println("Listener UDP address - ", ServerConn.LocalAddr().String())

	buf := make([]byte, 1024)
	for {
		n, addr, err := ServerConn.ReadFromUDP(buf)
		// If requester IP+port !equal our app's IP+port
		if addr.String() != myAddressWPort+":"+port {
			// if !stringInSlice(string(buf[0:n]), usefullIps) {
			go websocket_work.ConnectTo(string(buf[0:n]))
			// }
		}
		if err != nil {
			log.Println("Error: ", err)
		}
	}
}

// func stringInSlice(a string, list []string) bool {
// 	for _, b := range list {
// 		if b == a {
// 			return true
// 		}
// 	}
// 	return false
// }

// func gotUsefullIP(ip string, myport string) {
// 	for _, port := range AvailabalePortsTCP {
// 		if port != myport {
// 			go websocket_work.ConnectTo(ip + ":" + port)
// 		}
// 	}
// }
