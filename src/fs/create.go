package fs

import (
	"os"
)

// CreateDirectory - create dir
func CreateDirectory(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}

// CreateFile - create file
func CreateFile(path string) error {
	var _, err = os.Stat(path)
	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		defer file.Close()
		if err != nil {
			return err
		}
	}
	return nil
}
