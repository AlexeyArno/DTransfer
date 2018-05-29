package gui

import (
	"fmt"
	"log"
	"strings"

	"github.com/AlexeyArno/golang-files-transfer/src/utility"
	webview "github.com/zserge/webview"
)

var (
	currentWindow *webview.WebView
)

func HandleRPC(w webview.WebView, data string) {
	switch {
	case data == "opendir":
		answer := w.Dialog(webview.DialogTypeOpen, webview.DialogFlagDirectory, "Open directory", "")
		answer = strings.Replace(answer, `\`, "/", -1)
		s := fmt.Sprintf(`transData = {path: "%s"}`, answer)
		w.Eval(s)
	case data == "myIP":
		UpdateIP(w)

	}
}

func RegisterGUI(window *webview.WebView) {
	currentWindow = window
}

func UpdateIP(w webview.WebView) {
	IP, err := utility.MyIP()
	if err == nil {
		setIP(IP, w)
	} else {
		log.Println(err)
	}
}

func Log(logs string) {
	if currentWindow == nil {
		return
	}
	(*currentWindow).Dispatch(func() {
		(*currentWindow).Eval(fmt.Sprintf(`setLog("%s")`, logs))
	})
}
