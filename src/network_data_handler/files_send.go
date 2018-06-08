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
}

func BreakUpload() {
	if !uploadNow {
		return
	}
	breakChannel <- struct{}{}
}

func ContinueUpload() {
	if !uploadNow {
		return
	}
	continueChannel <- struct{}{}
}

func breakUploadNative() {
	// Clear all data from dir crawler in fs
	fs.Clear()
	nowBytesTransfered = 0
	totalPacketSendCount = 0
	packetCounter = 0
	packetTime = 0
	network_client.SendDoneRequest(totalPacketSendCount, RecieverIP)
	UploadCallback()
}

// StartSending - start scan and send data from selected dir = UploadPath
func StartSending() {
	breakChannel = make(chan (struct{}), 1)
	// Get amount bytes in dir
	totalDirSizeByte = fs.DirSizeByte(UploadPath)
	// Send reciever amount
	network_client.SendFullSize(totalDirSizeByte, RecieverIP)
	fs.SetPath(UploadPath)
	// Connnect to reciever
	conn, err := GetConnectionByIP(RecieverIP)
	if err != nil {
		log.Println("Start Sending: ", err)
		return
	} else {
		recieveConnection = conn
	}
	uploadNow = true
	defer func() { uploadNow = false }()
	workWithFiles()
}

func workWithFiles() {
	if !uploadNow {
		return
	}
	// Get next path
	path, isDir, err := fs.Next()
	if err != nil {
		// If all files visited
		if nowBytesTransfered == totalDirSizeByte {
			breakUploadNative()
		}
		return
	}
	if !isDir {
		currentPathLocker.Lock()
		currentPath = path
		currentPathLocker.Unlock()
		// Send reciever command create new file
		network_client.SendNewFileOffer(path, RecieverIP)
		err = readChunkAndSend(UploadPath + path)
		if err != nil {
			log.Println("StartSending 2: ", err)
			return
		}
	} else {
		// Send reciever command create new direcory
		network_client.SendNewDirectoryOffer(path, RecieverIP)
	}
	workWithFiles()
}

func readChunkAndSend(path string) error {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("ReadChunkAndSend 1:", err)
		return err
	}

	defer file.Close()

	data := make([]byte, DataSize)

	for {
		select {
		// If sender User click Stop
		// We're just wait click Continue
		case _ = <-stopChannel:
			<-continueChannel
		// User click break
		// We're break all transaction
		case _ = <-breakChannel:
			breakUploadNative()
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

	//Create 4 byte for amount of usefull bytes
	bs := make([]byte, 4)
	binary.LittleEndian.PutUint32(bs, uint32(n))

	finData = append(bs, (*data)...)

	begin := time.Now().UnixNano()

	err := recieveConnection.WriteMessage(websocket.BinaryMessage, finData)
	if err != nil {
		log.Println("Send files_send:", err)
		return err
	}

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
	totalPacketSendCount++
	packetCounter++

	return nil
}

// GetSpeed get middle speed for last 10 packets
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

// GetCurrentPath return current relative upload path
func GetCurrentPath() string {
	currentPathLocker.Lock()
	ret := currentPath
	currentPathLocker.Unlock()
	return ret
}

// GetProgress return [0:100] progress percent
func GetProgress() uint8 {
	return uint8((float64(nowBytesTransfered) / float64(totalDirSizeByte)) * 100)
}
