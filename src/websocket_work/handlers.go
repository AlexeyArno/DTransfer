package websocket_work

import (
	"log"
	"sync"

	"github.com/AlexeyArno/golang-files-transfer/src/utility"
	"golang.org/x/net/websocket"
)

var (
	machines       []*machine
	enteringData   = make(chan *websocket.Conn)
	leaving        = make(chan *websocket.Conn)
	machinesLocker = &sync.Mutex{}
)

func init() {
	go WebsocketService()
}

func DataHandler(ws *websocket.Conn) {

	log.Println("Client origin: ", ws.Request().RemoteAddr, "conn to data-channel")
	enteringData <- ws
	listenDataChannel(ws, 1)
}

// WebsocketService process all new and old websocket connections
func WebsocketService() {
	ip, err := utility.MyIP()
	if err != nil {
		return
	}

	for {
		select {

		case c := <-enteringData:
			if c == nil {
				log.Println("Data channel faile")
				return
			}
			m := machineIsHere(c)
			if m != nil {
				(*m).dataConn = c
				// log.Println("Data channel connected")
			} else {
				machinesLocker.Lock()
				nMachine := machine{dataConn: c}
				nMachine.IP = c.Config().Origin.Host
				if nMachine.IP == ip+":"+tcpPort {
					nMachine.IP = c.Request().Header.Get("Requester-Ip")
				}

				machines = append(machines, &nMachine)
				// log.Println("Data channel connected")
				machinesLocker.Unlock()
			}
		case c := <-leaving:
			m := machineIsHere(c)
			if m != nil {
				log.Println(m)
				if (*m).dataConn != nil {
					(*m).dataConn.Close()
				}
				deleteMachine(m)
				m = nil
				log.Println("Machine deleted")
			}
		}
	}
}

func machineIsHere(w *websocket.Conn) *machine {
	machinesLocker.Lock()
	defer machinesLocker.Unlock()
	for _, m := range machines {
		if (*m).IP == w.Config().Origin.Host {
			return m
		}
	}
	return nil
}

func deleteMachine(m *machine) {
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

}
