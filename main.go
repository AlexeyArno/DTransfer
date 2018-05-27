package main

import (
	"net/http"
	"strings"

	gui "github.com/AlexeyArno/golang-files-transfer/src/gui"
	"github.com/AlexeyArno/golang-files-transfer/src/network"
	"github.com/zserge/webview"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "Hello " + message
	w.Write([]byte(message))
}

func main() {
	// r := mux.NewRouter()
	// http.Handle("/", r)
	// go func() {
	// 	if err := http.ListenAndServe(":1002", nil); err != nil {
	// 		panic(err)
	// 	}
	// }()

	url := network.StartServer()
	w := webview.New(webview.Settings{
		Title: "DTransfer",
		URL:   url,
		// URL:                    `data:text/html,` + url.PathEscape(string(data)),
		Width:                  300,
		Height:                 400,
		ExternalInvokeCallback: gui.HandleRPC})

	// gui.ConnectGUIs(&w)
	defer w.Exit()
	w.Run()

}
