package network_data_handler

var TCPPort string
var UploadPath string
var DownloadPath string

var RecieverIP string
var SenderIP string

func RegisterTCPPort(IP string) {
	TCPPort = IP
}
