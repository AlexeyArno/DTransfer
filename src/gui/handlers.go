package gui

import (
	"fmt"

	webview "github.com/zserge/webview"
)

func setIP(ip string, w webview.WebView) {
	if currentWindow == nil {
		return
	}
	s := fmt.Sprintf(`setMyIP("%s")`, ip)
	w.Eval(s)
}
