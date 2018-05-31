package gui

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"
)

var tr = &http.Transport{
	MaxIdleConns:       10,
	IdleConnTimeout:    30 * time.Second,
	DisableCompression: true,
}

// startLoad: data = {"path": ..., "ip": 192...}
func startLoad(data map[string]string) {
	var Path, IP string
	if path, ok := data["path"]; ok {
		Path = path
	} else {
		log.Println("startLoad 1: 'path' doesnt exist")
		return
	}
	if ip, ok := data["ip"]; ok {
		IP = ip
	} else {
		log.Println("startLoad 2: 'ip' doesnt exist")
		return
	}
	splitPath := strings.Split(Path, "/")
	directoryName := splitPath[len(splitPath)-1]

	client := &http.Client{
		Transport: tr,
	}

	by := new(bytes.Buffer)
	dataFinal := map[string]string{"data": directoryName}
	json.NewEncoder(by).Encode(dataFinal)

	req, err := http.NewRequest("POST", "http://"+IP+"/info", by)
	if err != nil {
		log.Println("startLoad 3: ", err)
		return
	}
	req.Header.Add("Command-Type", "Offer")

	resp, err := client.Do(req)
	if err != nil {
		log.Println("startLoad 4: ", err)
		return
	}
	log.Println(resp.Body)

	// log.Println(Path, IP, directoryName)

}
