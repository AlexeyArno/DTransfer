package network_data_handler

import "log"

func init() {
	go startRecieve()
}

func startRecieve() {
	for {
		select {
		case pct := <-DownloadPacketSequence:
			recieve(pct.IP, &pct.Data)
		}
	}
}

func recieve(IP string, data *[]byte) {
	log.Println("Recieve from - ", IP, " ", len(*data), " bytes")
}
