package data_handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func parseCommand(req *http.Request) (map[string]interface{}, error) {
	var nCommand map[string]interface{}

	// switch req.Header.Get("Command-Type") {
	// case "Offer":
	// 	nCommand = command{}
	// 	break
	// case "NewDirectory":
	// 	nCommand = newDirectory{}
	// 	break
	// case "NewFile":
	// 	nCommand = newFile{}
	// 	break
	// }

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
	nCommand, err := parseCommand(req)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	switch req.Header.Get("Command-Type") {
	case "Offer":
		// nCommand = command{}
		offer(nCommand, req.Host)
		break
	case "NewDirectory":
		// nCommand = newDirectory{}
		break
	case "NewFile":
		// nCommand = newFile{}
		break
	}

	w.Header().Set("content-type", "application/json")
	w.Write([]byte(`{"yes": "no"}`))
	log.Println(nCommand["data"])
}
