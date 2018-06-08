package fs

import (
	"os"
	"path/filepath"
)

var dirSize uint64 = 0

// readSize return current dir's amount of bytes
func readSize(path string, file os.FileInfo, err error) error {
	if !file.IsDir() {
		dirSize += uint64(file.Size())
	}
	return nil
}

// DirSizeByte return whole dir's amount of bytes
func DirSizeByte(path string) uint64 {
	dirSize = 0
	filepath.Walk(path, readSize)
	return dirSize
}
