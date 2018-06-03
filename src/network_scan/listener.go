package network_scan

import (
	"log"
	"net"

	websocket_work "github.com/AlexeyArno/golang-files-transfer/src/websocket_work"
)

var (
	usefullIps []string
)

// Listen return hello
func Listen(myAddressWPort string, port string, ipFound func(string), messages *chan string) {

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
		// cleanIP := utility.GetCleanIPFromString(addr.String())
		// fmt.Println(cleanIP)
		if addr.String() != myAddressWPort+":"+port {
			if !stringInSlice(string(buf[0:n]), usefullIps) {
				// 	gotUsefullIP(cleanIP, port)
				go websocket_work.ConnectTo(string(buf[0:n]))
			}
			// go websocket_work.ConnectTo(ip + ":" + port)
		}
		if err != nil {
			log.Println("Error: ", err)
		}
	}
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func gotUsefullIP(ip string, myport string) {

	IPFound(ip)
	for _, port := range AvailabalePortsTCP {
		if port != myport {
			go websocket_work.ConnectTo(ip + ":" + port)
		}
	}
}
