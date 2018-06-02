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

func DialogInfo(body string) {
	if currentWindow == nil {
		return
	}

	(*currentWindow).Dispatch(func() {
		(*currentWindow).Dialog(webview.DialogTypeAlert, webview.DialogFlagInfo,
			"DTransfer", body)
	})
}

func GetOffer(dirName string, IP string) {
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

func ShippedOfferAccepted() {
	if currentWindow == nil {
		return
	}
	(*currentWindow).Dispatch(func() {
		s := fmt.Sprintf(`offerAccepted()`)
		(*currentWindow).Eval(s)
	})
}

func ShippedOfferCanceled(reason string) {
	if currentWindow == nil {
		return
	}
	(*currentWindow).Dispatch(func() {
		s := fmt.Sprintf(`offerCanceled("%s")`, reason)
		(*currentWindow).Eval(s)
	})
}
