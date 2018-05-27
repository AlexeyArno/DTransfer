package network

import "log"

func StartScanNetworks(CIDR_addresses []string) {
	var networksIP []IP

	var ip IP
	for _, addr := range CIDR_addresses {
		err := ip.FromString(addr)
		if err == nil {
			networksIP = append(networksIP, ip)
			ip.GetNetwork()
		} else {
			log.Println(err)
		}
	}

}
