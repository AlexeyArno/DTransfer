package network_client

import (
	"log"
	"strconv"

	"github.com/AlexeyArno/golang-files-transfer/src/utility"
)

var MyIP string
var myTCPPort string

func RegTCP(port string) {
	myTCPPort = port
}

func init() {
	ip, err := utility.MyIP()
	if err != nil {
		return
	}
	MyIP = ip
}

func SendOffer(directoryName, ip string) error {
	dataFinal := map[string]interface{}{"data": directoryName}
	headers := map[string]string{"Command-Type": "Offer", "Requester-IP": MyIP + ":" + myTCPPort}

	_, err := sendRequest(&dataFinal, &headers, ip)
	if err != nil {
		log.Println("Send Offer: ", err)
		return err
	}
	return nil
}

func SendAcceptOffer(ip string) error {
	headers := map[string]string{"Command-Type": "Accept-Offer",
		"Requester-IP": MyIP + ":" + myTCPPort}
	_, err := sendRequest(nil, &headers, ip)
	if err != nil {
		log.Println("SendAcceptOffer: ", err)
		return err
	}
	return nil
}

func SendCancelOffer(ip string) error {
	headers := map[string]string{"Command-Type": "Cancel-Offer",
		"Requester-IP": MyIP + ":" + myTCPPort}

	_, err := sendRequest(nil, &headers, ip)
	if err != nil {
		log.Println("SendCancelOffer : ", err)
		return err
	}
	return nil
}

func SendNewFileOffer(relativePath string, ip string) error {
	headers := map[string]string{"Command-Type": "File",
		"Requester-IP": MyIP + ":" + myTCPPort,
		"Path":         relativePath}

	_, err := sendRequest(nil, &headers, ip)
	if err != nil {
		log.Println("SendNewFileOffer : ", err)
		return err
	}
	return nil
}

func SendNewDirectoryOffer(relativePath string, ip string) error {
	headers := map[string]string{"Command-Type": "Dir",
		"Requester-IP": MyIP + ":" + myTCPPort,
		"Path":         relativePath}

	_, err := sendRequest(nil, &headers, ip)
	if err != nil {
		log.Println("SendNewDirectoryOffer : ", err)
		return err
	}
	return nil
}

func SendDoneRequest(packetCount uint64, ip string) error {
	headers := map[string]string{"Command-Type": "Done",
		"Requester-IP": MyIP + ":" + myTCPPort,
		"Packets":      strconv.FormatUint(packetCount, 10)}

	_, err := sendRequest(nil, &headers, ip)
	if err != nil {
		log.Println("SendDoneRequest : ", err)
		return err
	}
	return nil
}

func SendBrokeRequest(ip string) error {
	headers := map[string]string{"Command-Type": "Broke",
		"Requester-IP": MyIP + ":" + myTCPPort}

	_, err := sendRequest(nil, &headers, ip)
	if err != nil {
		log.Println("SendBrokeRequest : ", err)
		return err
	}
	return nil
}

func SendFullSize(size uint64, ip string) error {
	headers := map[string]string{"Command-Type": "Size",
		"Requester-IP": MyIP + ":" + myTCPPort,
		"Size":         strconv.FormatUint(size, 10)}

	_, err := sendRequest(nil, &headers, ip)
	if err != nil {
		log.Println("SendDoneRequest : ", err)
		return err
	}
	return nil
}
