package gui

import (
	"fmt"

	webview "github.com/zserge/webview"
)

// RegisterGUI - set current GUI Window
func RegisterGUI(window *webview.WebView) {
	currentWindow = window
}

// Alert - set Alert label in HTML
func Alert(reason string) {
	if currentWindow == nil {
		return
	}
	(*currentWindow).Dispatch(func() {
		(*currentWindow).Eval(fmt.Sprintf(`alertPage("%s")`, reason))
	})
}

// DialogInfo - call Dialog window
func DialogInfo(body string) {
	if currentWindow == nil {
		return
	}

	(*currentWindow).Dispatch(func() {
		(*currentWindow).Dialog(webview.DialogTypeAlert, webview.DialogFlagInfo,
			"DTransfer", body)
	})
}

// GetOffer - add new transaction offer to sequence in HTML
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

// ShippedOfferAccepted - make some HTML objects manipulation
func ShippedOfferAccepted() {
	if currentWindow == nil {
		return
	}
	(*currentWindow).Dispatch(func() {
		s := fmt.Sprintf(`offerAccepted()`)
		(*currentWindow).Eval(s)
	})
}

// ShippedOfferCanceled - make some HTML objects manipulation
func ShippedOfferCanceled(reason string) {
	if currentWindow == nil {
		return
	}
	(*currentWindow).Dispatch(func() {
		s := fmt.Sprintf(`offerCanceled("%s")`, reason)
		(*currentWindow).Eval(s)
	})
}

// UploadDone - make some HTML objects manipulation
func UploadDone() {
	if currentWindow == nil {
		return
	}
	(*currentWindow).Dispatch(func() {
		s := fmt.Sprintf(`uploadDone()`)
		(*currentWindow).Eval(s)
	})
}

// DownloadDone - make some HTML objects manipulation
func DownloadDone() {
	if currentWindow == nil {
		return
	}
	(*currentWindow).Dispatch(func() {
		s := fmt.Sprintf(`breakDownload()`)
		(*currentWindow).Eval(s)
	})
}
