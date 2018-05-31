package data_handler

import gui "github.com/AlexeyArno/golang-files-transfer/src/gui"

func offer(c map[string]interface{}, IP string) {
	gui.Offer(c["data"].(string), IP)
}
