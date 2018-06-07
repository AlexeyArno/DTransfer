package gui_js_api

import (
	"fmt"
	"log"
	"strings"

	"github.com/AlexeyArno/golang-files-transfer/src/network_data_handler"
	"github.com/AlexeyArno/golang-files-transfer/src/utility"
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
	IPs := network_data_handler.GetConnectedIPs()
	log.Println(IPs)
	s := fmt.Sprintf(`redrawIPs(["%s"])`, strings.Join(IPs, ","))

	(*w).Eval(s)
}

func GetDashboardData(w *webview.WebView) {
	speed := network_data_handler.GetSpeed()
	path := network_data_handler.GetCurrentPath()
	progress := network_data_handler.GetProgress()
	// log.Println("Progress: ", progress, "%")
	s := fmt.Sprintf(`refreshDashboard("%s", %d, %d)`, path, speed, progress)
	(*w).Eval(s)
}

func GetDashboardDataDownload(w *webview.WebView) {
	path := network_data_handler.GetCurrentRelativePath()
	progress := network_data_handler.GetDownloadProgress()
	// log.Println("Progress Download: ", progress, "%")
	s := fmt.Sprintf(`refreshDashboardDownload("%s", %d)`, path, progress)
	(*w).Eval(s)
}
