package main

import (
	"net/http"
	"strings"

	"github.com/zserge/webview"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "Hello " + message
	w.Write([]byte(message))
}

func main() {
	http.HandleFunc("/", sayHello)
	go func() {
		if err := http.ListenAndServe(":8080", nil); err != nil {
			panic(err)
		}
	}()

	webview.Open("Hello", "http://127.0.0.1:8080", 400, 300, false)
}
