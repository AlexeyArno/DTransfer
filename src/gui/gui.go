package gui

import (
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
	}
}
