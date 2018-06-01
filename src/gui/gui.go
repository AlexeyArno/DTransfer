package gui

import (
	"encoding/json"
	"log"
	"strings"

	gui_js_api "github.com/AlexeyArno/golang-files-transfer/src/gui_js_api"
	webview "github.com/zserge/webview"
)

var (
	currentWindow *webview.WebView
)

func HandleRPC(w webview.WebView, data string) {
	switch {
	case data == "opendir":
		gui_js_api.OpenDir(&w)
	case data == "myIP":
		gui_js_api.UpdateIP(&w)
	case data == "getConnectedIPs":
		gui_js_api.GetConnectedIPs(&w)
	case strings.HasPrefix(data, "acceptTransitOffer:"):
		IP := strings.TrimPrefix(data, "acceptTransitOffer:")
		log.Println("Accept Transfer: ", IP)
	case strings.HasPrefix(data, "cancelTransitOffer:"):
		IP := strings.TrimPrefix(data, "cancelTransitOffer:")
		log.Println("Cancel Transfer: ", IP)
	case strings.HasPrefix(data, "startLoad:"):
		data := strings.TrimPrefix(data, "startLoad:")
		log.Println(data)
		var dat map[string]string

		if err := json.Unmarshal([]byte(data), &dat); err != nil {
			log.Println("Handle RPC: ", err)
			return
		}

		go startLoad(dat)
	}
}
