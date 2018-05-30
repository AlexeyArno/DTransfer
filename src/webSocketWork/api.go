package webSocketWork

import "log"

// GetConnectedIPs return connected machines IPs
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
