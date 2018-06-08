package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/AlexeyArno/golang-files-transfer/src/gui"
	"github.com/AlexeyArno/golang-files-transfer/src/network_data_handler"
	"github.com/AlexeyArno/golang-files-transfer/src/network_scan"
	network_server "github.com/AlexeyArno/golang-files-transfer/src/network_server"
	"github.com/AlexeyArno/golang-files-transfer/src/websocket_work"
	"github.com/zserge/webview"
)

func main() {
	log.SetOutput(ioutil.Discard)
	port, err := network_scan.ReservTCPPort()
	if err != nil {
		panic(err)
	}

	network_data_handler.RegisterTCPPort(port)
	log.Println("Listen TCP port: ", port)

	go network_scan.SetupSanner()

	data, err := gui.Asset("data/index.html")
	if err != nil {
		// panic(err)
		return
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write(data)
	})
	http.HandleFunc("/info", network_server.InfoHandler)
	http.HandleFunc("/websocket_data", websocket_work.DataHandler)

	go func(port string) {
		if err := http.ListenAndServe(":"+port, nil); err != nil {
			// panic(err)
			return
		}
	}(port)

	w := webview.New(webview.Settings{
		Title:                  "DTransfer",
		URL:                    "http://127.0.0.1:" + port,
		Width:                  300,
		Height:                 420,
		ExternalInvokeCallback: gui.HandleRPC})
	gui.RegisterGUI(&w)
	defer w.Exit()
	w.Run()
}
