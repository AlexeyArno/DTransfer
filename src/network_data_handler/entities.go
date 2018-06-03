package network_data_handler

import (
	"errors"
	"log"

	"github.com/gorilla/websocket"
)

type Packet struct {
	IP   string
	Data []byte
}

type Machine struct {
	dataConn *websocket.Conn
	IP       string
}

func AddMachine(IP string, conn *websocket.Conn) {
	machinesLocker.Lock()
	nMachine := Machine{dataConn: conn}
	nMachine.IP = IP
	machines = append(machines, &nMachine)
	machinesLocker.Unlock()
}

func GetConnectionByIP(IP string) (*websocket.Conn, error) {
	machinesLocker.Lock()
	defer machinesLocker.Unlock()
	for _, m := range machines {
		if (*m).IP == IP {
			return (*m).dataConn, nil
		}
	}
	return nil, errors.New("Machine didn't find")
}

func GetIPByConnection(c *websocket.Conn) (string, error) {
	machinesLocker.Lock()
	defer machinesLocker.Unlock()
	for _, m := range machines {
		if (*m).dataConn == c {
			return (*m).IP, nil
		}
	}
	return "", errors.New("Machine didn't find")
}

func MachineIsHere(w *websocket.Conn) bool {
	machinesLocker.Lock()
	defer machinesLocker.Unlock()
	for _, m := range machines {
		if (*m).dataConn == w {
			return true
		}
	}
	return false
}

func DeleteMachine(w *websocket.Conn) {
	machinesLocker.Lock()
	for i, machine := range machines {
		if (*machine).dataConn == w {
			machines[i] = machines[len(machines)-1]
			machines[len(machines)-1] = nil
			machines = machines[:len(machines)-1]
			break
		}
	}
	machinesLocker.Unlock()

}

func GetConnectedIPs() []string {
	var finish []string
	machinesLocker.Lock()
	log.Println(machines)
	for _, j := range machines {
		log.Println(j)
		if j.dataConn != nil {
			finish = append(finish, j.IP)
		}
	}
	machinesLocker.Unlock()
	return finish
}
