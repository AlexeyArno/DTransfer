package network_server

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/AlexeyArno/golang-files-transfer/src/fs"

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

	ip := req.Header.Get("Requester-IP")
	if len(ip) == 0 {
		http.Error(w, "Requester-IP expected", 500)
		return
	}

	switch req.Header.Get("Command-Type") {
	case "Offer":
		nCommand, err := parseCommand(req)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		offer(nCommand, ip)
		break
	case "Accept-Offer":
		// go start
		network_data_handler.RecieverIP = ip
		log.Println(ip, " accepted your offer")
		gui.ShippedOfferAccepted()
		go network_data_handler.StartSending(gui.UploadDone)

		break
	case "Cancel-Offer":
		log.Println(ip, " canceled your offer")
		network_data_handler.RecieverIP = ""
		gui.ShippedOfferCanceled(ip + " disline your offer")
		break
	case "File":
		path := req.Header.Get("Path")
		if len(path) == 0 {
			http.Error(w, "Path expected", 500)
			return
		}
		if ip == network_data_handler.SenderIP {
			fs.CreateFile(network_data_handler.GetDownloadPath() + path)
			network_data_handler.ChangeRelativeDownloadPath(path)
			w.Write([]byte(`1`))
			return
		} else {
			http.Error(w, "Rights aren't", 500)
		}
		break
	case "Dir":
		path := req.Header.Get("Path")
		if len(path) == 0 {
			http.Error(w, "Path expected", 500)
			return
		}
		if ip == network_data_handler.SenderIP {
			fs.CreateDirectory(network_data_handler.GetDownloadPath() + path)
			w.Write([]byte(`1`))
			return
		} else {
			http.Error(w, "Rights aren't", 500)
		}
		break
	case "Done":
		if ip == network_data_handler.SenderIP {
			pockets := req.Header.Get("Packets")
			if len(pockets) == 0 {
				http.Error(w, "Pockets expected", 500)
				return
			}
			i, err := strconv.ParseUint(pockets, 10, 64)
			if err != nil {
				panic(err)
			}
			network_data_handler.Done(i)
			w.Write([]byte(`1`))
			return
		} else {
			http.Error(w, "Rights aren't", 500)
		}
		break
	case "Size":
		log.Println("Size Got!!!!!!!!")
		if ip == network_data_handler.SenderIP {
			size := req.Header.Get("Size")
			log.Println("My Size is ", size)
			if len(size) == 0 {
				http.Error(w, "Size expected", 500)
				return
			}
			i, err := strconv.ParseUint(size, 10, 64)
			if err != nil {
				panic(err)
			}
			network_data_handler.SetFullSize(i)
			w.Write([]byte(`1`))
			return
		}
	}

	// w.Header().Set("content-type", "application/json")
	w.Write([]byte(`1`))
	// log.Println(nCommand["data"])
}
