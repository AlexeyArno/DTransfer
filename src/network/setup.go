package network

//StartServer start server
func StartServer() string {
	// // data, err := gui.Asset("data/index.html")
	// if err != nil {
	// 	panic(err)
	// }

	StartScanNetworks()

	// ln, err := net.Listen("tcp", "127.0.0.1:1235")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// go func() {
	// 	defer ln.Close()
	// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 		w.Write(data)
	// 	})

	// 	log.Fatal(http.Serve(ln, nil))
	// }()

	// return "http://" + ln.Addr().String()
	return "192.0.0.0"
}
