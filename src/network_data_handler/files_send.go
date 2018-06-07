package network_data_handler

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/AlexeyArno/golang-files-transfer/src/network_client"

	"github.com/AlexeyArno/golang-files-transfer/src/fs"
	"github.com/gorilla/websocket"
)

var recieveConnection *websocket.Conn
var uploaded chan (struct{})

var stopChannel chan (struct{}) = make(chan (struct{}), 1)
var breakChannel chan (struct{}) = make(chan (struct{}), 1)
var continueChannel chan (struct{}) = make(chan (struct{}), 1)

var uploadNow bool

var packetCounter uint8

var packetSum uint32 = 0
var packetTime uint32 = 0

var totalPacketSendCount uint64 = 0

var totalDirSizeByte uint64
var nowBytesTransfered uint64

func StopUpload() {
	if !uploadNow {
		return
	}

	stopChannel <- struct{}{}
	// log.Println("StopUpload")
}

func BreakUpload() {
	if !uploadNow {
		return
	}

	breakChannel <- struct{}{}
	// log.Println("BreakUpload")
}

func ContinueUpload() {
	if !uploadNow {
		return
	}
	continueChannel <- struct{}{}
	// log.Println("ContinueUpload")
}

func init() {
	// go initSending()
}

// StartSending very well
func StartSending(uploadDoneCallback func()) {
	breakChannel = make(chan (struct{}), 1)
	totalDirSizeByte = fs.DirSizeByte(UploadPath)
	network_client.SendFullSize(totalDirSizeByte, RecieverIP)
	// log.Println("RecieverIP: ", RecieverIP)
	fs.SetPath(UploadPath)
	conn, err := GetConnectionByIP(RecieverIP)
	if err != nil {
		log.Println("Start Sending: ", err)
		return
	} else {
		recieveConnection = conn
	}
	uploadNow = true
	defer func() { uploadNow = false }()
	workWithFiles(uploadDoneCallback)

	// go getFiles()
}

func workWithFiles(uploadDoneCallback func()) {
	// log.Println("Work with files")
	if !uploadNow {
		return
	}
	path, isDir, err := fs.Next()
	if err != nil {
		log.Println("StartSending: ", err)
		log.Println("Transfered: ", nowBytesTransfered, "/", totalDirSizeByte)
		if nowBytesTransfered == totalDirSizeByte {
			fs.Clear()
			nowBytesTransfered = 0
			totalPacketSendCount = 0
			packetCounter = 0
			packetTime = 0
			network_client.SendDoneRequest(totalPacketSendCount, RecieverIP)
			uploadDoneCallback()
		}
		return
	}
	if !isDir {
		currentPathLocker.Lock()
		currentPath = path
		currentPathLocker.Unlock()
		network_client.SendNewFileOffer(path, RecieverIP)
		err = readChunkAndSend(UploadPath+path, uploadDoneCallback)
		if err != nil {
			log.Println("StartSending 2: ", err)
			return
		}
	} else {
		network_client.SendNewDirectoryOffer(path, RecieverIP)
	}
	workWithFiles(uploadDoneCallback)
}

func readChunkAndSend(path string, uploadDoneCallback func()) error {

	// log.Println("Read Chunk ", path)

	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	data := make([]byte, DataSize)

	for {
		select {
		case _ = <-stopChannel:
			<-continueChannel
		case _ = <-breakChannel:
			fs.Clear()
			network_client.SendBrokeRequest(RecieverIP)
			nowBytesTransfered = 0
			totalPacketSendCount = 0
			packetCounter = 0
			packetTime = 0
			return errors.New("Sender break upload")
		default:
			n, err := file.Read(data)
			if err == io.EOF { // если конец файла
				return nil // выходим из цикла
			}
			err = send(n, &data)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func send(n int, data *[]byte) error {
	finData := make([]byte, BufferSize)
	bs := make([]byte, 4)
	binary.LittleEndian.PutUint32(bs, uint32(n))
	finData = append(bs, (*data)...)

	begin := time.Now().UnixNano()

	err := recieveConnection.WriteMessage(websocket.BinaryMessage, finData)
	if err != nil {
		// log.Println("Send file_send:", err)
		return err
	}
	// bytsCount := binary.LittleEndian.Uint32(finData[:4])

	packetSum += uint32(n)
	packetTime += uint32((time.Now().UnixNano() - begin) / int64(time.Millisecond))
	if packetCounter >= 10 {
		lastPacketTimeLocker.Lock()
		last10PacketTime = packetTime
		last10PacketSize = packetSum
		lastPacketTimeLocker.Unlock()
		packetCounter = 0
	}
	nowBytesTransfered += uint64(n)
	// log.Println("Sended  ", nowBytesTransfered, "/", totalDirSizeByte, " KB")
	totalPacketSendCount++
	packetCounter++

	return nil
}

func GetSpeed() uint32 {
	lastPacketTimeLocker.Lock()
	if last10PacketTime == 0 {
		last10PacketTime = 1
	}
	ret := (last10PacketSize / last10PacketTime) * 1000
	lastPacketTimeLocker.Unlock()
	return ret
	// return uint64(finish)
}

func GetCurrentPath() string {
	currentPathLocker.Lock()
	ret := currentPath
	currentPathLocker.Unlock()
	return ret
}

func GetProgress() uint8 {
	return uint8((float64(nowBytesTransfered) / float64(totalDirSizeByte)) * 100)
}
