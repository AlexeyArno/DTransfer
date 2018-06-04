package network_data_handler

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"time"

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

func StopUpload() {
	if !uploadNow {
		return
	}

	stopChannel <- struct{}{}
	log.Println("StopUpload")
}

func BreakUpload() {
	if !uploadNow {
		return
	}

	breakChannel <- struct{}{}
	log.Println("BreakUpload")
}

func ContinueUpload() {
	if !uploadNow {
		return
	}
	continueChannel <- struct{}{}
	log.Println("ContinueUpload")
}

func init() {
	// go initSending()
}

// StartSending very well
func StartSending() {

	log.Println("RecieverIP: ", RecieverIP)
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
	workWithFiles()
	// go getFiles()
}

func workWithFiles() {
	log.Println("Work with files")
	if !uploadNow {
		return
	}
	path, idDir, err := fs.Next()
	if err != nil {
		log.Println("StartSending: ", err)
		return
	}
	if !idDir {
		currentPathLocker.Lock()
		currentPath = path
		currentPathLocker.Unlock()
		err = readChunkAndSend(path)
		if err != nil {
			log.Println("StartSending 2: ", err)
			return
		}
	}
	time.Sleep(time.Second)
	workWithFiles()
}

func readChunkAndSend(path string) error {

	log.Println("Read Chunk ", path)

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
			log.Println("Im breaking")
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
		log.Println("Send file_send:", err)
		return err
	}
	// log.Println("Send Chunk ", n)
	packetSum += uint32(n)
	packetTime += uint32((time.Now().UnixNano() - begin) / int64(time.Millisecond))
	if packetCounter >= 10 {
		lastPacketTimeLocker.Lock()
		last10PacketTime = packetTime
		last10PacketSize = packetSum
		lastPacketTimeLocker.Unlock()
		packetCounter = 0
	}
	packetCounter++

	return nil
}
