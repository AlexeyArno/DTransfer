package network_server

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/AlexeyArno/golang-files-transfer/src/gui"
	"github.com/AlexeyArno/golang-files-transfer/src/network_data_handler"
)

func parseCommand(req *http.Request) (map[string]interface{}, error) {
	var nCommand map[string]interface{}

	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nCommand, err
	}

	err = json.Unmarshal(b, &nCommand)
	if err != nil {
		return nCommand, err
	}
	return nCommand, nil
}

func InfoHandler(w http.ResponseWriter, req *http.Request) {

	switch req.Header.Get("Command-Type") {
	case "Offer":
		nCommand, err := parseCommand(req)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		ip := req.Header.Get("Requester-IP")
		if len(ip) == 0 {
			http.Error(w, "Requester-IP expected", 500)
		}
		offer(nCommand, ip)
		break
	case "Accept-Offer":
		// go start
		ip := req.Header.Get("Requester-IP")
		if len(ip) == 0 {
			http.Error(w, "Requester-IP expected", 500)
		}
		network_data_handler.RecieverIP = ip
		log.Println(ip, " accepted your offer")
		gui.ShippedOfferAccepted()
		go network_data_handler.StartSending()

		break
	case "Cancel-Offer":
		ip := req.Header.Get("Requester-IP")
		if len(ip) == 0 {
			http.Error(w, "Requester-IP expected", 500)
		}
		log.Println(ip, " canceled your offer")
		network_data_handler.RecieverIP = ""
		gui.ShippedOfferCanceled(ip + " disline your offer")
		break
	case "NewDirectory":
		// nCommand = newDirectory{}
		break
	case "NewFile":
		// nCommand = newFile{}
		break
	}

	// w.Header().Set("content-type", "application/json")
	w.Write([]byte(`1`))
	// log.Println(nCommand["data"])
}
