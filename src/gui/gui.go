package gui

import (
	"fmt"
	"strings"

	webview "github.com/zserge/webview"
)

// ConnectGUIs to the web page
func ConnectGUIs(w *webview.WebView) {
	progress, err := Asset("data/progress.css")
	if err != nil {
		fmt.Println(err)
	}
	style, err := Asset("data/style.css")
	if err != nil {
		fmt.Println(err)
	}

	(*w).Dispatch(func() {
		(*w).InjectCSS(string(style))
		(*w).InjectCSS(string(progress))
	})
}

func HandleRPC(w webview.WebView, data string) {
	switch {
	case data == "opendir":
		answer := w.Dialog(webview.DialogTypeOpen, webview.DialogFlagDirectory, "Open directory", "")
		answer = strings.Replace(answer, `\`, "/", -1)
		s := fmt.Sprintf(`transData = {path: "%s"}`, answer)
		fmt.Println(s)
		w.Eval(s)

	}
}
