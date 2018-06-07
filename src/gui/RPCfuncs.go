package gui

import (
	"net/http"
	"strings"
	"time"

	"github.com/AlexeyArno/golang-files-transfer/src/network_client"
)

var tr = &http.Transport{
	MaxIdleConns:       10,
	IdleConnTimeout:    30 * time.Second,
	DisableCompression: true,
}

// startLoad: data = {"path": ..., "ip": 192...}
func sendOffer(data map[string]string) {
	var Path, IP string
	if path, ok := data["path"]; ok {
		Path = path
	} else {
		// log.Println("startLoad 1: 'path' doesnt exist")
		return
	}
	if ip, ok := data["ip"]; ok {
		IP = ip
	} else {
		// log.Println("startLoad 2: 'ip' doesnt exist")
		return
	}
	splitPath := strings.Split(Path, "/")
	directoryName := splitPath[len(splitPath)-1]

	err := network_client.SendOffer(directoryName, IP)
	if err != nil {
		// log.Println("sendOffer gui: ", err)
	}

}

func sendAcceptOffer(IP string) {
	err := network_client.SendAcceptOffer(IP)
	if err != nil {
		// log.Println("sendAcceptOffer gui: ", err)
	}

}

func sendCancelOffer(IP string) {
	err := network_client.SendCancelOffer(IP)
	if err != nil {
		// log.Println("sendCancelOffer gui: ", err)
	}
}
