package webSocketWork

import (
	"log"
	"sync"

	gui "github.com/AlexeyArno/golang-files-transfer/src/gui"
	"golang.org/x/net/websocket"
)

var (
	machines       []*machine
	enteringInfo   = make(chan *websocket.Conn)
	enteringData   = make(chan *websocket.Conn)
	leaving        = make(chan *websocket.Conn)
	machinesLocker = &sync.Mutex{}
)

func init() {
	go WebsocketService()
}

func DataHandler(ws *websocket.Conn) {
	gui.Log("Client conn to data-channel: " + ws.RemoteAddr().String())
	log.Println("Client conn to data-channel: ", ws.RemoteAddr().String())
	enteringData <- ws
	go socketListener(ws, 0)
}

func InfoHandler(ws *websocket.Conn) {
	gui.Log("Client conn to info-channel: " + ws.RemoteAddr().String())
	log.Println("Client conn to info-channel: ", ws.RemoteAddr().String())
	enteringInfo <- ws
	go socketListener(ws, 1)
}

// WebsocketService process all new and old websocket connections
func WebsocketService() {
	for {
		select {
		case c := <-enteringInfo:
			m := machineIsHere(c)
			if m != nil {
				(*m).infoConn = c
				(*m).IP = c.RemoteAddr().String()
			} else {
				machinesLocker.Lock()
				nMachine := machine{infoConn: c}
				nMachine.IP = c.RemoteAddr().String()

				machines = append(machines, &nMachine)
				machinesLocker.Unlock()
			}
		case c := <-enteringData:
			m := machineIsHere(c)
			if m != nil {
				(*m).dataConn = c
				(*m).IP = c.RemoteAddr().String()
			} else {
				machinesLocker.Lock()
				nMachine := machine{dataConn: c}
				nMachine.IP = c.RemoteAddr().String()

				machines = append(machines, &nMachine)
				machinesLocker.Unlock()
			}
		case c := <-leaving:
			m := machineIsHere(c)
			if m != nil {
				if (*m).dataConn != nil {
					(*m).dataConn.Close()
				}
				if (*m).infoConn != nil {
					(*m).infoConn.Close()
				}
				machinesLocker.Lock()
				for i, machine := range machines {
					if (*machine).IP == (*m).IP {
						machines[i] = machines[len(machines)-1]
						machines[len(machines)-1] = nil
						machines = machines[:len(machines)-1]
						break
					}
				}
				machinesLocker.Unlock()
				m = nil
			}
		}
	}
}

func machineIsHere(w *websocket.Conn) *machine {
	machinesLocker.Lock()
	defer machinesLocker.Unlock()
	for _, m := range machines {
		if (*m).IP == w.RemoteAddr().String() {
			return m
		}
	}
	return nil
}
