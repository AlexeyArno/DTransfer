package gui_js_api

import (
	"fmt"
	"log"
	"strings"

	"github.com/AlexeyArno/golang-files-transfer/src/utility"
	"github.com/AlexeyArno/golang-files-transfer/src/webSocketWork"
	webview "github.com/zserge/webview"
)

func OpenDir(w *webview.WebView) {
	answer := (*w).Dialog(webview.DialogTypeOpen, webview.DialogFlagDirectory, "Open directory", "")
	answer = strings.Replace(answer, `\`, "/", -1)
	s := fmt.Sprintf(`transData = {path: "%s"}`, answer)
	(*w).Eval(s)
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
	IPs := webSocketWork.GetConnectedIPs()
	log.Println(IPs)
	s := fmt.Sprintf(`redrawIPs(["%s"])`, strings.Join(IPs, ","))

	(*w).Eval(s)
}
