package network_data_handler

import (
	"bytes"
	"encoding/binary"
	"log"
	"os"
	"sync"
)

var downoloadLocker = &sync.Mutex{}
var donwloadNow bool
var currentFile *os.File

var currentRelativeFilesPath string
var pathLocker = &sync.Mutex{}

var bytesWritten uint64
var needSizeBytes uint64

var needPockets uint64
var PocketsLocker = &sync.Mutex{}

var pocketsNow uint64

var downloadSizeBytes uint64

var breakChannelDownload chan (struct{}) = make(chan (struct{}), 1)

var packetCounterDownload uint8

var DonwloadDoneCallback func()

func BreakDownload() {
	if !donwloadNow {
		return
	}

	breakChannelDownload <- struct{}{}
	log.Println("BreakDownload")
}

func BeginDownload() {
	downoloadLocker.Lock()
	donwloadNow = true
	downoloadLocker.Unlock()
}

func SetFullSize(size uint64) {
	needSizeBytes = size
}

func GetDownloadProgress() uint8 {
	return uint8((float64(bytesWritten) / float64(needSizeBytes)) * 100)
}

func Done(pockets uint64) {
	PocketsLocker.Lock()
	needPockets = pockets
	PocketsLocker.Unlock()
	if pocketsNow == pockets {
		// log.Println("Close 3")
		currentFile.Close()
		DonwloadDoneCallback()
	}
}

func GetDownloadPath() string {
	pathLocker.Lock()
	path := DownloadPath
	pathLocker.Unlock()
	return path
}

func ChangeRelativeDownloadPath(newPath string) {
	pathLocker.Lock()
	currentRelativeFilesPath = newPath
	file, err := os.OpenFile(DownloadPath+currentRelativeFilesPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Println("recieve: ", err)
		return
	}
	currentFile = file
	bytesWritten = 0
	pathLocker.Unlock()
}

func init() {
	go startRecieve()
}

func startRecieve() {
	for {
		select {
		case pct := <-DownloadPacketSequence:
			if pct.IP == SenderIP && donwloadNow {
				recieve(pct.IP, pct.Data)
				if pocketsNow == needPockets {
					log.Println("Close 2")
					currentFile.Close()
					DonwloadDoneCallback()
				}
			}
		case _ = <-breakChannelDownload:
			log.Println("Close 3")
			donwloadNow = false
			currentFile.Close()
			DonwloadDoneCallback()
		}
	}
}

func recieve(IP string, data []byte) {
	var bytsCount int32
	buf := bytes.NewBuffer((data)[:4])
	binary.Read(buf, binary.LittleEndian, &bytsCount)

	_, err := currentFile.Write(data[4 : bytsCount+4])
	if err != nil {
		log.Println("Recive last: ", err)
		return
	}

	bytesWritten += uint64(bytsCount)
	pocketsNow++
}
