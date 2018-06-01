package gui

import (
	"fmt"

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
		s := fmt.Sprintf(`offersPush({'text':"%s", 'IP':"%s"})`,
			IP+" offer transfer directory - '"+dirName+"'",
			IP)

		(*currentWindow).Eval(s)
	})

}
