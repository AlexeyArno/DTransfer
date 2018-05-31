package gui

import (
	webview "github.com/zserge/webview"
)

func RegisterGUI(window *webview.WebView) {
	currentWindow = window
}

// func Log(logs string) {
// 	if currentWindow == nil {
// 		return
// 	}
// 	(*currentWindow).Dispatch(func() {
// 		(*currentWindow).Eval(fmt.Sprintf(`setLog("%s")`, logs))
// 	})
// }

func Offer(dirName string, IP string) {
	if currentWindow == nil {
		return
	}
	(*currentWindow).Dispatch(func() {
		(*currentWindow).Dialog(
			webview.DialogTypeAlert, 0,
			IP+" offer transwer direcotory - "+dirName+" to you", "Ok")
	})

}
