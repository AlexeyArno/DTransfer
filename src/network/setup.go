package network

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"

	"github.com/AlexeyArno/golang-files-transfer/src/gui"
)

//StartServer start server
func StartServer() string {
	data, err := gui.Asset("data/index.html")
	if err != nil {
		panic(err)
	}
	ln, err := net.Listen("tcp", "127.0.0.1:1235")
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		defer ln.Close()
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write(data)
		})

		log.Fatal(http.Serve(ln, nil))
	}()

	StartScanNetworks(GetIPs())

	return "http://" + ln.Addr().String()
}

//GetIPs get IPs
func GetIPs() []string {
	var ips []string
	ifaces, err := net.Interfaces()
	if err != nil {
		log.Print(fmt.Errorf("localAddresses: %v\n", err.Error()))
		return ips
	}
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			log.Print(fmt.Errorf("localAddresses: %v\n", err.Error()))
			continue
		}
		for _, a := range addrs {

			if is_ip(a.String()) {
				ips = append(ips, a.String())
			}
		}
	}
	return ips
}

func is_ip(host string) bool {
	parts := strings.Split(host, ".")

	if len(parts) < 4 {
		return false
	}

	for _, x := range parts {
		if i, err := strconv.Atoi(x); err == nil {
			if i < 0 || i > 255 {
				return false
			}
		} else {
			parts2 := strings.Split(x, "/")
			if len(parts2) != 2 {
				return false
			}
		}
	}
	return true
}

func setupIP(CIDR_address string) {
	var ip IP
	err := ip.FromString(CIDR_address)
	if err != nil {
		log.Println(err)
	}
	ip.Print()
}
