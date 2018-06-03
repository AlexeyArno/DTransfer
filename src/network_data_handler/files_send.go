package network_data_handler

import (
	"github.com/AlexeyArno/golang-files-transfer/src/fs"
)

func StartSending() {
	fs.SetPath(UploadPath)
	go getFiles()
}

func getFiles() {
	// path, _, err := fs.Next()
	// if err == nil {
	// 	log.Println(path)
	// 	time.Sleep(time.Millisecond * 100)
	// 	getFiles()
	// } else {
	// 	log.Println("GetFiles: ", err)
	// 	files, dirs := fs.GetStat()
	// 	log.Println("Stat: files - ", files, " dirs - ", dirs)
	// }

}

func startSend() {

}
