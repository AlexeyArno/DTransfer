package network_data_handler

import (
	"log"

	"golang.org/x/net/websocket"
)

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

func MachineIsHere(w *websocket.Conn) bool {
	machinesLocker.Lock()
	defer machinesLocker.Unlock()
	for _, m := range machines {
		if (*m).IP == w.Config().Origin.Host || (*m).IP == w.Request().Header.Get("Partner-Ip") {
			return true
		}
	}
	return false
}

func DeleteMachine(w *websocket.Conn) {
	machinesLocker.Lock()
	for i, machine := range machines {
		if machine.IP == w.Config().Origin.Host || machine.IP == w.Request().Header.Get("Partner-Ip") {
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
