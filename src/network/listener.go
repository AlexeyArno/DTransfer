package network

import (
	"log"
	"net"
)

func Listen() {
	ServerAddr, err := net.ResolveUDPAddr("udp", ":10001")
	if err != nil {
		log.Println(err)
		return
	}

	/* Now listen at selected port */
	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	if err != nil {
		log.Println(err)
		return
	}
	defer ServerConn.Close()

	buf := make([]byte, 1024)

	for {
		_, addr, err := ServerConn.ReadFromUDP(buf)
		log.Println("Received ", " from ", addr)

		if err != nil {
			log.Println("Error: ", err)
		}
	}
}
