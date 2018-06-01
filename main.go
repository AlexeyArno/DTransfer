package main

import (
	"log"
	"net/http"

	data_handler "github.com/AlexeyArno/golang-files-transfer/src/data_handler"
	"github.com/AlexeyArno/golang-files-transfer/src/gui"
	"github.com/AlexeyArno/golang-files-transfer/src/network"
	websocket_work "github.com/AlexeyArno/golang-files-transfer/src/websocket_work"
	"github.com/zserge/webview"
	"golang.org/x/net/websocket"
)

var close = make(chan struct{})

func main() {
	data, err := gui.Asset("data/index.html")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write(data)
	})
	http.HandleFunc("/info", data_handler.InfoHandler)
	http.Handle("/websocket_data", websocket.Handler(websocket_work.DataHandler))

	port, err := network.ReservTCPPort()
	if err != nil {
		panic(err)
	}
	log.Println("Listen TCP port: ", port)
	websocket_work.RegisterTCPPort(port)

	go func(port string) {
		if err := http.ListenAndServe(":"+port, nil); err != nil {
			panic(err)
		}
	}(port)

	// close <- struct{}{}
	// network.StartServer()
	w := webview.New(webview.Settings{
		Title:                  "DTransfer",
		URL:                    "http://127.0.0.1:" + port,
		Width:                  300,
		Height:                 420,
		ExternalInvokeCallback: gui.HandleRPC})
	gui.RegisterGUI(&w)
	defer w.Exit()
	w.Run()
	<-close

}
