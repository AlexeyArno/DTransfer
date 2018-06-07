package gui

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/AlexeyArno/golang-files-transfer/src/network_data_handler"

	gui_js_api "github.com/AlexeyArno/golang-files-transfer/src/gui_js_api"
	webview "github.com/zserge/webview"
)

var (
	currentWindow *webview.WebView
)

func HandleRPC(w webview.WebView, data string) {
	switch {
	case data == "opendir":
		network_data_handler.UploadPath = gui_js_api.OpenDir(&w)
	case data == "myIP":
		gui_js_api.UpdateIP(&w)
	case data == "getConnectedIPs":
		gui_js_api.GetConnectedIPs(&w)
	case data == "stopUpload":
		go network_data_handler.StopUpload()
	case data == "breakUpload":
		go network_data_handler.BreakUpload()
	case data == "continueUpload":
		go network_data_handler.ContinueUpload()
	case data == "getDashboardData":
		gui_js_api.GetDashboardData(currentWindow)
	case data == "getDashboardDataDownload":
		gui_js_api.GetDashboardDataDownload(&w)
	case strings.HasPrefix(data, "acceptTransitOffer:"):
		IP := strings.TrimPrefix(data, "acceptTransitOffer:")
		log.Println("I accept to ", IP)
		network_data_handler.DownloadPath = gui_js_api.OpenDir(&w)
		network_data_handler.SenderIP = IP
		network_data_handler.DonwloadDoneCallback = DownloadDone
		network_data_handler.BeginDownload()
		sendAcceptOffer(IP)
		break
	case strings.HasPrefix(data, "cancelTransitOffer:"):
		IP := strings.TrimPrefix(data, "cancelTransitOffer:")
		log.Println("I dismiss to ", IP)
		network_data_handler.SenderIP = ""
		network_data_handler.DownloadPath = ""
		sendCancelOffer(IP)
		break
	case strings.HasPrefix(data, "dialoginfo:"):
		body := strings.TrimPrefix(data, "dialoginfo:")
		DialogInfo(body)
	case strings.HasPrefix(data, "startLoad:"):
		data := strings.TrimPrefix(data, "startLoad:")
		// log.Println(data)
		var dat map[string]string

		if err := json.Unmarshal([]byte(data), &dat); err != nil {
			log.Println("Handle RPC: ", err)
			return
		}

		if ip, ok := dat["ip"]; ok {
			network_data_handler.RecieverIP = ip
		} else {
			log.Println("handelRPC startLoad 2: 'ip' doesnt exist")
			return
		}

		go sendOffer(dat)
	}
}
