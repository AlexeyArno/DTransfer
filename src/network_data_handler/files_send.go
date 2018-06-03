package network_data_handler

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/AlexeyArno/golang-files-transfer/src/fs"
	"github.com/gorilla/websocket"
)

var recieveConnection *websocket.Conn
var uploaded chan (struct{})

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
	workWithFiles()
	// go getFiles()
}

func workWithFiles() {
	log.Println("Work with files")
	path, idDir, err := fs.Next()
	if err != nil {
		log.Println("StartSending: ", err)
		return
	}
	if !idDir {
		err = readChunkAndSend(path)
		if err != nil {
			log.Println("StartSending 2: ", err)
			return
		}
	}
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
		n, err := file.Read(data)
		if err == io.EOF { // если конец файла
			break // выходим из цикла
		}
		err = send(n, &data)
		if err != nil {
			return err
		}
	}
	// time.Sle
	return nil
}

func send(n int, data *[]byte) error {
	log.Println("Send chunk")
	finData := make([]byte, BufferSize)
	bs := make([]byte, 4)
	binary.LittleEndian.PutUint32(bs, uint32(n))
	log.Println("Counter Buffer: ", bs)
	finData = append(bs, (*data)...)

	log.Println("Send Chunk ", finData[:4], " ", finData[5:8], " ", n)
	err := recieveConnection.WriteMessage(websocket.BinaryMessage, finData)
	if err != nil {
		log.Println("Send file_send:", err)
		return err
	}

	return nil
}
