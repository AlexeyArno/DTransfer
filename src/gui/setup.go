package gui

import "github.com/AlexeyArno/golang-files-transfer/src/network_data_handler"

func init() {
	network_data_handler.UploadCallback = UploadDone
	network_data_handler.DonwloadDoneCallback = DownloadDone
}
