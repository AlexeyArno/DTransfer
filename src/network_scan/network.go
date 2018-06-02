package network_scan

import (
	"errors"
	"net"
)

// ReservUDPPort check every default UDP port for available
func ReservUDPPort() (string, error) {
	for _, port := range AvailabalePortsUDP {
		ServerAddr, err := net.ResolveUDPAddr("udp", ":"+port)
		if err == nil {
			ServerConn, err := net.ListenUDP("udp", ServerAddr)
			if err != nil {
				continue
			} else {
				ServerConn.Close()
				return port, nil
			}
		}
	}
	return "", errors.New("Didn't find available port")
}

// ReservTCPPort check every default TCP port for available
func ReservTCPPort() (string, error) {
	for _, port := range AvailabalePortsTCP {
		ln, err := net.Listen("tcp", ":"+port)
		if err != nil {
			continue
		} else {
			ln.Close()
			return port, nil
		}
	}
	return "", errors.New("Didn't find available port")
}
