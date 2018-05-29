package utility

import (
	"errors"
	"log"
	"net"
	"os"
	"strings"
)

func CheckError(err error) {
	if err != nil {
		log.Println("Error: ", err)
		os.Exit(0)
	}
}

// GetCleanIPFromString transform "192.168.0.1:8080"->"192.168.0.1"
func GetCleanIPFromString(myAddress string) string {
	return strings.Split(myAddress, ":")[0]
}

func MyIP() (string, error) {
	myAddress := ""
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				myAddress = ipnet.IP.String()
			}
		}
	}
	if len(myAddress) == 0 {
		return "", errors.New("Cannot find my ip")
	} else {
		return myAddress, nil
	}

}
