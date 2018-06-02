package network_client

import (
	"log"

	"github.com/AlexeyArno/golang-files-transfer/src/network_data_handler"
	"github.com/AlexeyArno/golang-files-transfer/src/utility"
)

var MyIP string

func init() {
	ip, err := utility.MyIP()
	if err != nil {
		return
	}
	MyIP = ip
}

func SendOffer(directoryName, ip string) error {

	dataFinal := map[string]interface{}{"data": directoryName}
	headers := map[string]string{"Command-Type": "Offer", "Requester-IP": MyIP + ":" + network_data_handler.TCPPort}

	_, err := sendRequest(&dataFinal, &headers, ip)
	if err != nil {
		log.Println("Send Offer: ", err)
		return err
	}
	return nil
}

func SendAcceptOffer(ip string) error {
	headers := map[string]string{"Command-Type": "Accept-Offer",
		"Requester-IP": MyIP + ":" + network_data_handler.TCPPort}
	_, err := sendRequest(nil, &headers, ip)
	if err != nil {
		log.Println("SendAcceptOffer: ", err)
		return err
	}
	return nil
}

func SendCancelOffer(ip string) error {

	headers := map[string]string{"Command-Type": "Cancel-Offer",
		"Requester-IP": MyIP + ":" + network_data_handler.TCPPort}

	_, err := sendRequest(nil, &headers, ip)
	if err != nil {
		log.Println("SendCancelOffer : ", err)
		return err
	}
	return nil
}
