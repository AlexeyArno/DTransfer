package network

import (
	"log"
	"net"
	"time"
)

func BroadcastToNetwork() {

	log.Println("Setup Broadcast")

	ServerAddr, err := net.ResolveUDPAddr("udp", "255.255.255.255:10001")
	if err != nil {
		log.Println(err)
		return
	}

	addrs, err := net.InterfaceAddrs()
	var myAddres string
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				myAddres = ipnet.IP.String() + ":0"
			}
		}
	}

	log.Println("My Address: ", myAddres)

	LocalAddr, err := net.ResolveUDPAddr("udp", myAddres)
	if err != nil {
		log.Println(err)
		return
	}

	Conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
	if err != nil {
		log.Println(err)
		return
	}

	defer Conn.Close()
	for i := 0; i < 4; i++ {
		_, err := Conn.Write([]byte{})
		if err != nil {
			log.Println(err)
		}
		time.Sleep(time.Second * 1)
	}

}
