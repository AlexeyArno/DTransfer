package network_data_handler

import (
	"log"
)

func init() {
	go startRecieve()
}

func startRecieve() {
	for {
		select {
		case pct := <-DownloadPacketSequence:
			if pct.IP == SenderIP {
				recieve(pct.IP, &pct.Data)
			} else {
				log.Println(pct.IP, "!=", SenderIP)
			}
		}
	}
}

func recieve(IP string, data *[]byte) {
	log.Println("Received: ", (*data)[:4], " ", (*data)[5:8])

	// log.Println("Recieve from - ", IP, " ", len(*data), " bytes")
}
