package network_server

import gui "github.com/AlexeyArno/golang-files-transfer/src/gui"

func offer(c map[string]interface{}, IP string) {
	gui.GetOffer(c["data"].(string), IP)
}
