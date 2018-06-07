package fs

import (
	"log"
	"os"
)

func CreateDirectory(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, 0755)
		if err != nil {
			panic(err)
		} else {
			log.Println("Create Dir:", path)
		}
	}
}

func CreateFile(path string) {
	var _, err = os.Stat(path)

	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if err != nil {
			panic(err)
		} else {
			log.Println("Create File:", path)
		}
		defer file.Close()
	}

}
