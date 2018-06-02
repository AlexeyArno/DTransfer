package gui_js_api

import (
	"fmt"
	"log"
	"strings"

	"github.com/AlexeyArno/golang-files-transfer/src/utility"
	websocket_work "github.com/AlexeyArno/golang-files-transfer/src/websocket_work"
	webview "github.com/zserge/webview"
)

func OpenDir(w *webview.WebView) string {
	answer := (*w).Dialog(webview.DialogTypeOpen, webview.DialogFlagDirectory, "Open directory", "")
	answer = strings.Replace(answer, `\`, "/", -1)
	// network_data_handler.UploadPath = answer
	s := fmt.Sprintf(`transData = {path: "%s"}`, answer)
	(*w).Eval(s)
	return answer
}

func UpdateIP(w *webview.WebView) {
	IP, err := utility.MyIP()
	if err == nil {
		s := fmt.Sprintf(`setMyIP("%s")`, IP)
		(*w).Eval(s)
	} else {
		log.Println(err)
	}
}

func GetConnectedIPs(w *webview.WebView) {
	IPs := websocket_work.GetConnectedIPs()
	log.Println(IPs)
	s := fmt.Sprintf(`redrawIPs(["%s"])`, strings.Join(IPs, ","))

	(*w).Eval(s)
}
