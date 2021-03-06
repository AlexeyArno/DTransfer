package network_client

import (
	"strconv"

	"github.com/AlexeyArno/golang-files-transfer/src/utility"
)

var MyIP string
var myTCPPort string

// RegTCP - set TCP port for package to request
// that _reciever know on which port send answer
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

// SendOffer - send to Reciever offer for begin transaction
func SendOffer(directoryName, ip string) error {
	dataFinal := map[string]interface{}{"data": directoryName}
	headers := map[string]string{"Command-Type": "Offer", "Requester-IP": MyIP + ":" + myTCPPort}

	_, err := sendRequest(&dataFinal, &headers, ip)
	if err != nil {
		// log.Println("Send Offer: ", err)
		return err
	}
	return nil
}

// SendAcceptOffer - send positive answer to Sender request
func SendAcceptOffer(ip string) error {
	headers := map[string]string{"Command-Type": "Accept-Offer",
		"Requester-IP": MyIP + ":" + myTCPPort}
	_, err := sendRequest(nil, &headers, ip)
	if err != nil {
		// log.Println("SendAcceptOffer: ", err)
		return err
	}
	return nil
}

// SendCancelOffer - send negative answer to Sender requester
func SendCancelOffer(ip string) error {
	headers := map[string]string{"Command-Type": "Cancel-Offer",
		"Requester-IP": MyIP + ":" + myTCPPort}

	_, err := sendRequest(nil, &headers, ip)
	if err != nil {
		// log.Println("SendCancelOffer : ", err)
		return err
	}
	return nil
}

// SendNewFileOffer - send to Reciever request create new file
func SendNewFileOffer(relativePath string, ip string) error {
	headers := map[string]string{"Command-Type": "File",
		"Requester-IP": MyIP + ":" + myTCPPort,
		"Path":         relativePath}

	_, err := sendRequest(nil, &headers, ip)
	if err != nil {
		// log.Println("SendNewFileOffer : ", err)
		return err
	}
	return nil
}

//  SendNewDirectoryOffer - send to Reciever request create new dir
func SendNewDirectoryOffer(relativePath string, ip string) error {
	headers := map[string]string{"Command-Type": "Dir",
		"Requester-IP": MyIP + ":" + myTCPPort,
		"Path":         relativePath}

	_, err := sendRequest(nil, &headers, ip)
	if err != nil {
		// log.Println("SendNewDirectoryOffer : ", err)
		return err
	}
	return nil
}

// SendDoneRequest - send to Reciever message that sender done read and send
func SendDoneRequest(packetCount uint64, ip string) error {
	headers := map[string]string{"Command-Type": "Done",
		"Requester-IP": MyIP + ":" + myTCPPort,
		"Packets":      strconv.FormatUint(packetCount, 10)}

	_, err := sendRequest(nil, &headers, ip)
	if err != nil {
		// log.Println("SendDoneRequest : ", err)
		return err
	}
	return nil
}

// SendBrokeRequest - send offer for break send to Sender
func SendBrokeRequest(ip string) error {
	headers := map[string]string{"Command-Type": "Broke",
		"Requester-IP": MyIP + ":" + myTCPPort}

	_, err := sendRequest(nil, &headers, ip)
	if err != nil {
		// log.Println("SendBrokeRequest : ", err)
		return err
	}
	return nil
}

// SendFullSize - send dir size to Reciever
func SendFullSize(size uint64, ip string) error {
	headers := map[string]string{"Command-Type": "Size",
		"Requester-IP": MyIP + ":" + myTCPPort,
		"Size":         strconv.FormatUint(size, 10)}

	_, err := sendRequest(nil, &headers, ip)
	if err != nil {
		// log.Println("SendDoneRequest : ", err)
		return err
	}
	return nil
}
