package network

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type IP struct {
	body [4]byte
	mask byte
}

type Network struct {
	begin IP
	last  IP
}

func (ip *IP) FromString(CIDR_address string) error {
	err := errors.New("CIDR address is incorrect, expected 192.168.0.1/24, but got" + CIDR_address)
	parts := strings.Split(CIDR_address, ".")

	if len(parts) < 4 {
		return err
	}
	var checker uint8 = 0
	for _, x := range parts {
		if i, err := strconv.Atoi(x); err == nil {
			if i < 0 || i > 255 {
				return err
			}
			ip.body[checker] = byte(i)
			checker++
		} else {
			parts2 := strings.Split(x, "/")
			if len(parts2) != 2 {
				return err
			}
			if i, err := strconv.Atoi(parts2[0]); err == nil {
				if i < 0 || i > 255 {
					return err
				}
				ip.body[checker] = byte(i)
				checker++
			}
			if i, err := strconv.Atoi(parts2[1]); err == nil {
				ip.mask = byte(i)
			}
		}
	}
	return nil

}

func (ip *IP) Print() {
	fmt.Print(ip.body[0])
	for _, b := range ip.body {
		fmt.Print(".")
		fmt.Print(b)
	}
	fmt.Print("/", ip.mask)
	fmt.Println("")
}

func (ip *IP) GetNetwork() {
	// var netw Network

	var mask [4]byte
	j := int(ip.mask) / 8
	i := 0
	for ; i < j; i++ {
		mask[i] = 0xff
	}

	k := ip.mask % 8
	switch k {
	case 1:
		mask[i] = 0x80
	case 2:
		mask[i] = 0xC0
	case 3:
		mask[i] = 0xE0
	case 4:
		mask[i] = 0xF0
	case 5:
		mask[i] = 0xF8
	case 6:
		mask[i] = 0xFC
	case 7:
		mask[i] = 0xFE
	}

	log.Println(ip.mask, " - ", mask)

}
