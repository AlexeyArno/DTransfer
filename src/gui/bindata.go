// Code generated by go-bindata.
// sources:
// data/index.html
// DO NOT EDIT!

package gui

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _dataIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x3d\x7f\x8f\x9b\xb8\xb6\xff\x5f\xe9\x7e\x07\x97\x6a\x95\x4c\x0b\x0e\x3f\x03\xa4\x93\x4a\xbd\xb3\x5d\xed\xe8\x75\x77\x47\x6a\xdf\xd5\xbb\x5a\xad\x46\x1e\x70\x12\x54\x82\x59\x70\x32\xc9\x56\xfd\xee\x4f\xb6\x81\x81\x04\x1b\x32\xbb\x7d\xda\x97\x6a\x12\x02\xc7\xe7\x1c\x9f\xdf\x3e\x86\xf4\xfa\x45\x4c\x22\x7a\xcc\x31\xd8\xd0\x6d\xfa\xf6\x9f\xff\xb8\xae\x3e\x01\x00\xe0\x7a\x83\x51\x5c\x1d\xd7\xaf\xeb\x2d\xa6\x08\x6c\x28\xcd\x0d\xfc\xfb\x2e\xd9\x2f\xb5\xff\x31\xfe\xfb\x9d\x71\x43\xb6\x39\xa2\xc9\x43\x8a\x35\x10\x91\x8c\xe2\x8c\x2e\xb5\xdb\xf7\xcb\xf0\xcd\xed\xfb\xa5\x65\xb2\xf7\xf7\xf1\x1a\xeb\xd1\xa6\x20\x5b\xbc\xb4\xb4\xd9\x29\xde\x92\x1e\x53\x7c\x72\x92\xbd\xe0\x63\x81\xf2\x1c\x17\xe0\xcb\xf9\x35\xf6\x7a\x4c\x62\xba\x59\x00\xcb\x37\xf3\xc3\x1b\x30\x7b\x05\x3e\x62\x0a\xe8\x06\x83\x32\xf9\x03\x03\xb2\xe2\xc7\x79\x41\xd6\x05\x2e\x4b\xf0\x80\x0a\xf0\x6a\xd6\x8f\x69\x83\x93\xf5\x86\xd6\xa8\xfa\x61\xfe\x30\x92\x2c\xc6\x87\x05\xb0\x6c\x53\x0e\x95\x93\x32\xa1\x09\xc9\x16\x00\x3d\x94\x24\xdd\x51\xcc\x39\x7b\x9f\xa1\x87\x14\x83\x28\x4d\xf2\x3c\xc9\xd6\x52\x46\x18\xc0\x02\x14\x38\xa2\x53\x33\x3f\xe8\xc0\xf2\xda\x1f\xbe\x97\x1f\xae\x38\xbe\x1f\x93\x18\x83\x0d\x4a\x57\x17\x4d\xf3\x6b\xff\xe9\x13\xc9\x95\x35\x4e\x9c\xe2\x2d\xce\x68\x09\xe8\x06\x51\xb0\x45\x9f\x31\xd8\xe5\xe3\xa8\xc1\x28\x29\xa2\x14\xcb\x34\x37\x4e\x96\xb5\x7e\x1d\x39\x48\xa3\x38\x05\xcc\x03\x29\x62\x5c\x2c\x80\x65\xe6\x07\x50\x92\x34\x89\xc1\xcb\x20\xb8\x31\x3d\x5f\x39\xc0\x28\x50\x9c\xec\xca\x05\x08\x9d\x4b\xb4\x3d\x4e\xaf\x4c\x93\x8d\x5a\x4d\xa6\x55\xa9\x6a\x68\x81\xb2\x9a\x0a\x3f\x5e\x91\x62\x0b\x80\x09\x9d\xf2\x8d\x54\xd3\xf5\x38\x06\xbb\x00\x05\xa1\x88\xe2\xa9\x19\xe3\xf5\x95\x7c\xcc\x89\x75\x5c\xcf\xfa\xfc\x52\xea\xac\x0f\x24\x3e\x4a\xb4\xfd\x80\xa2\xcf\xeb\x82\xec\xb2\x78\x01\x5e\xda\x3f\x38\xbe\xfb\x4e\xad\x71\x47\x61\x14\x2b\x92\x51\x63\x85\xb6\x49\x7a\x5c\x80\x77\x45\x82\x52\x1d\xfc\x88\xd3\x3d\xa6\x49\x84\x74\x50\xa2\xac\x34\x4a\x5c\x24\x2b\xd5\x78\x66\xe6\xcc\xf8\xa4\x54\x28\x3e\x50\x03\xa5\xc9\x3a\x5b\x80\x08\x67\x14\x17\x32\x0b\x40\x71\x9c\x64\x6b\x83\x92\x5c\x58\x98\xcc\x00\x48\x4a\x8a\xc5\x4b\x6c\xb2\x7f\x32\x98\x5d\x51\x92\x62\x11\xe3\x15\xda\xa5\x54\x02\x44\xf6\xb8\x58\xa5\xe4\x71\x01\x36\x49\x1c\xe3\x4c\x02\x66\x3c\xe2\x87\xcf\x09\x35\x28\xd9\x45\x1b\x23\x42\x69\x4a\x76\x74\x01\x32\x92\x89\x80\x94\xfc\xf2\x11\x7c\x44\x2b\x54\x24\x52\x7b\xa8\x51\xec\x4a\x5c\x18\x25\x4e\x71\xd4\x46\x30\x34\xf8\x33\xcb\x26\x92\xb1\xff\x45\xb2\xdf\x77\xb8\x20\x05\xf8\xf1\xd3\x4f\x1f\xe4\x38\xb6\xe4\x0f\x09\x86\x1f\x92\x02\xaf\xc8\x41\x3a\x54\x0c\x2f\x25\xa3\x6f\x99\x42\x33\x4c\xc1\xfb\x43\x9e\x92\x02\x17\x33\x96\xa1\x94\xc8\xd8\xab\x1f\xd9\xcf\x24\x33\xf2\x02\xaf\x92\x03\x8e\xc1\x1e\x17\x65\x42\x32\x9d\xe9\xb2\xc0\x19\x4d\x8f\x6a\x94\x7d\xaf\x72\x97\xe7\xa4\xa0\x38\x06\x0f\x47\x70\xc3\xb3\x26\x40\x59\x0c\x7e\xc9\x71\x81\x7a\x99\xec\x89\xeb\xb3\x57\xe0\x13\xc9\x7b\xa1\x5f\x6e\x8f\xb7\x77\x12\x57\xad\xcd\xf9\x81\x50\x4a\xb6\x0b\x60\x4b\x2d\x7a\x14\x17\x50\x41\x8a\xe4\x28\x4a\xe8\x71\x01\x4c\xe8\xf5\x90\x38\xc7\xd6\x3b\xc9\xbb\x3a\x13\xf5\xa7\xa1\x3a\x51\xdd\xb3\xca\x04\x25\x19\x2e\x06\x6a\x09\x4f\xee\xc2\x5b\x54\xac\x13\x16\xea\x77\x94\x0c\xa5\x23\x05\x1a\x16\x9a\x49\xbe\x30\x44\xca\x52\x45\xf0\xa7\xec\x52\xe0\x14\xd1\x64\x8f\x95\xe0\x9d\xa4\x2a\x87\x14\xd3\xa8\x02\x56\xa0\x62\xb3\x1d\xb7\x0b\x1c\xf7\xa3\xfc\xfa\xcf\x7f\x9c\x9f\xec\x31\xba\x46\x13\x39\x49\x98\xfb\x49\xf4\xb0\x4a\x09\xa2\x0b\x50\x30\x39\x0e\x47\x70\xdb\x1b\x0a\xb8\x2b\xfe\x52\x61\x7a\xac\x74\xe6\x9b\xb2\xb8\xdc\xa3\x07\x49\xd6\x20\xf9\x62\x3e\x54\xcf\x58\x72\x9f\x1a\x9f\x76\xb8\x7c\x16\x56\xff\xf4\x47\x2a\x09\xa6\x04\xc5\xb8\xb8\x2f\x37\x28\x26\x8f\x03\x7e\xa1\xa8\xaf\x46\xf8\xc5\xe8\x3a\x69\x4c\x3d\xd7\xb2\x73\x09\x44\x8a\x57\x74\x11\xc8\xcb\x36\xa6\x27\xb5\x87\xb6\x4d\xff\x21\xdd\x29\x1c\xef\xa4\x58\xf4\xcc\xef\xc6\x97\xa1\xc5\xfa\x01\x4d\x4d\x9d\xff\x83\x16\xab\x01\x47\xaa\xae\x47\xc5\xe0\x65\xb4\x21\xa4\xc4\x4a\xbf\xe2\x66\x33\x2a\x72\xf7\x98\x0b\xab\x25\x07\xd2\x86\xaa\x0a\x1a\x47\x24\x23\x8f\xf7\xab\x24\xed\x9f\x46\xe5\xd3\x91\xcb\xfe\xf5\x8a\xb9\x0a\x1f\xcc\x00\xde\x8c\x4b\x51\xbd\x59\xe5\xfb\xa4\xc0\x11\x29\x8e\x92\xac\x12\xb3\xcb\x94\x14\xc7\xfb\xb1\x32\x51\x98\xf3\xb3\x9d\xaa\x72\xcd\x7a\x05\x35\x2e\x21\x0b\x23\xc9\xee\x9b\x29\xa8\xe3\x30\x17\xe4\x73\x75\x89\xb2\x64\x8b\x28\x3e\xbf\x22\xa1\xd9\x5e\xe9\xa0\x34\x05\x26\xb4\xca\xa1\xf2\x76\xdc\x90\x3e\x49\xa0\x88\x0d\xbb\x7f\xd8\x51\x4a\x32\x09\x47\x55\x45\x0e\xaa\x9c\xa5\x54\xc4\x16\x1d\x8c\xaa\xff\x21\x81\x8b\x93\x32\x4f\xd1\x71\x01\x92\x2c\x4d\x32\x59\x08\x3c\xaf\x32\x15\xb9\x22\x66\x66\x8a\xc4\xfc\x25\xb0\x83\x73\x5f\x6c\xd8\x9a\x42\x26\x01\xee\x72\x3c\xfd\x8f\x43\xcd\x3c\xe2\x5e\xb0\xaf\x44\x59\x65\xe6\xb0\x0f\x2f\x38\x29\x53\x14\x19\xb3\x02\xe3\x11\x1f\x18\xfd\xf9\xb0\x87\xcd\x97\x82\xc3\xfb\x24\x1f\xc8\x7b\x8a\x0a\xe3\x24\xf2\xdb\x50\x01\xda\x5e\x01\x2b\xc2\xd7\xc9\x7c\xfe\x4c\x24\x7d\xf9\x80\xd7\x89\xcc\xae\x67\xaf\xda\xe5\xb7\x33\xae\x58\x54\x68\x61\x9c\xcc\x39\x47\xf7\xe5\x7e\xad\xf7\x29\x84\x92\x5c\x76\xed\xa1\xc0\xe8\xb3\xec\xa2\xa8\x62\x24\x13\x3d\x71\x39\xe3\x21\x25\xd1\x67\x75\x3c\x55\x14\x06\x75\x75\x32\xbc\x48\x50\x84\xed\xde\xd6\x8c\xe3\x88\xde\xcc\xd8\xe5\x1d\x4c\xb6\x68\x8d\xd5\xc1\xeb\xe2\xf9\x8c\xf4\x1c\x85\xa2\x84\x2e\x94\x2a\x1c\x52\xd4\xf8\x28\xa6\xd6\xfc\xb9\x8c\xed\xaa\xff\x35\x0e\xb9\x68\x60\x7e\x90\x4e\xa8\x7d\xdd\x1e\xcd\xc3\x05\x2c\x48\xc2\xf6\x7e\x40\xeb\x7d\x19\x11\x60\x54\x62\x23\xc9\x0c\xb2\x93\xe5\xa6\x71\x99\xae\x15\x34\xc2\x01\x9b\x52\x54\x3c\x4d\x97\x6f\x6c\x74\xeb\x4e\x5c\x99\xaf\x1a\x0e\xad\xe7\x60\x16\xdf\x64\x6d\x4c\x72\x30\xc4\x5a\x49\x99\x97\x5b\x32\xf2\xc7\x99\x1a\x4a\x71\xa1\x4e\x96\xfd\xe9\x17\x5c\xb2\x62\x7c\x6a\x4d\x0c\xe5\xd2\x0b\x8b\x78\x98\x12\x99\x53\xb3\x92\xa8\xa6\xeb\x8e\x5b\xf5\x4a\x0a\x4e\xd0\x98\x8d\x6a\xc3\x60\x38\xf8\x76\xb2\x9e\x04\x64\xb8\xc1\xda\xa7\xc4\x02\xb3\x70\x74\x9f\xe4\xa5\x32\x1c\x2b\x92\x68\x2d\x29\x05\xc8\xb8\x0a\xf2\xa2\x7e\x85\x2b\x25\xc6\x6b\x10\x45\x83\xe5\xd2\x88\x31\xbf\x58\x92\xe3\x3c\xdd\x84\xc1\xe5\x98\x95\x9e\xfe\x14\x44\x46\xba\x80\x08\x1f\xf7\x49\x34\x90\x8c\x9f\xaf\xfd\xde\x36\xc0\x4b\xb2\x5a\x0d\x75\x34\x1d\x85\xc3\x8f\xa8\x67\x3a\x3d\xc0\x76\xbb\xc2\x91\x6d\x59\x8d\xf4\xe6\x7a\x79\x2c\x2a\x5d\x39\x8f\xa3\x5b\x47\x02\x91\x22\xca\x30\x63\x57\x64\xa5\xce\x3e\x1b\xc9\xc7\x24\xce\x3e\x1b\xe3\x2a\xb9\x57\xec\x86\x3d\xa3\x24\xf5\xa5\x4d\xa5\x11\xd1\xb5\x13\xd5\x47\x2f\x90\x60\x35\x0f\x55\xa9\x71\xc9\x4c\x46\xc7\xa4\xda\x6c\x07\xd3\x54\x51\xd9\xee\xf0\xbe\x6f\xb7\xdf\x66\x7b\x9e\x5e\xff\x99\xd0\x93\x99\xf1\xc9\xea\x4e\xd1\x4c\x1c\x9b\x81\x95\x51\xe5\x22\x19\x5d\xba\xda\xed\x28\x53\x19\x54\xdb\x35\x8e\x99\x1f\xf8\x1f\x33\x1a\x7e\xd0\xf1\x7f\xbf\x57\x72\x72\x8f\x38\x62\x49\x62\x3c\x8b\x30\xb6\x1b\xe8\xc0\xf2\x6d\x1d\x38\x16\x23\x74\x75\x09\x99\x8c\xf4\x53\x19\x19\x51\x3b\x92\x92\x65\xf2\x4b\xec\x7e\x84\x83\x8e\xd8\xf2\x78\x9e\x0f\x9f\x9f\xe9\xee\xe9\x5f\xcf\x5a\x77\xf9\x5c\xb3\x90\xd5\xdd\xd7\xef\x8e\xbf\x8e\x93\x3d\x88\x52\x54\x96\x4b\x6d\x7b\xbc\xbd\xd3\xde\xfe\x87\xec\x0a\x70\x7b\xb7\x00\xd7\x65\x8e\x32\x90\xc4\xf5\x05\xcb\xf6\xa1\x09\x4d\x68\x5d\xcf\xd8\x95\xb7\xd7\xb3\x38\xd9\x9f\xde\x48\xd0\x42\xd7\xd9\x88\xd0\x06\xc1\xcf\x37\xf5\xb4\x9e\x1b\x12\xda\x23\xaa\x1b\x89\x34\x10\x23\x8a\x0c\x94\x25\xdb\xa5\xf6\x80\x4a\x0c\x9a\x0b\x8c\xf9\xfa\x4b\x0f\xb2\x53\x84\x62\xf5\x77\x8e\x8f\x79\xa6\x40\x26\x20\xee\xf9\x89\xde\x19\x5d\x82\x96\xdb\x46\x07\xaf\x38\x23\x45\x2c\x3d\xcf\x88\x31\x34\xa7\xdb\x71\xb2\x49\x9b\xdf\x8d\x43\xde\x73\x4e\xae\xc0\x6e\xff\x7c\x48\x79\x67\x2d\x6b\x19\xaf\xdf\xd7\x00\x2d\x8b\x9c\x9c\x8d\x9e\xbc\x65\x8b\xb7\xca\x32\x2f\x94\x1c\x3d\xe6\x78\xa9\x89\xe8\x50\xa9\x83\xa3\xd7\x6a\x5e\x3b\x8d\x55\x19\x9f\x37\x7c\xcc\x8b\xcb\x05\x2b\x95\xe8\x84\xc9\x71\x32\x20\xc7\x56\x67\x56\x6a\xe2\x5c\x28\x1f\x71\x16\x03\x4a\x16\xa0\x12\x52\x75\x51\x36\x84\x63\xe4\xc2\x68\x9a\xaa\x32\xfc\x62\x04\xc9\x99\x90\xc0\x1e\xa5\x3b\xbc\xd4\xac\xd0\x86\xd6\x3c\x80\x26\xb4\xb5\xb7\x4f\x5f\xac\xeb\x99\x80\x93\xb1\x3a\x13\xd4\x54\x5e\xc5\x78\x7a\xaa\xf9\x55\x4c\x5d\x97\xfb\x75\x2d\xa7\x56\x31\xaf\x81\xc3\x36\xcd\xca\xa5\xb6\xa1\x34\x5f\xcc\x66\x8f\x8f\x8f\xf0\xd1\x81\xa4\x58\xcf\x6c\xd3\x34\x67\xe5\x7e\x5d\x81\x2c\x0e\x69\x92\x7d\xee\x03\xb4\xc2\x30\x9c\xf1\xab\x5a\x7d\x7b\xc8\x52\xb3\xa0\x25\xec\xe7\x06\xe5\xe8\xde\xd2\xc0\x61\xa9\x99\xf9\x41\x03\xc7\xea\x73\x9f\xe0\xc7\x7f\x11\x76\x16\x98\xc0\x0e\xa0\x3d\xf7\xaa\x0f\x0d\xf0\x18\xbe\xd4\x30\xbf\xa7\xd0\x68\xe5\xcf\x0c\x3f\x82\x33\xf8\x37\x9c\xc3\x45\x99\xa3\x08\x33\xbf\xc7\x25\x2e\xf6\x58\x13\x85\xd6\x52\xf3\x2c\x9b\xd1\x13\x79\xaa\xfe\xaa\x52\x1f\x93\xd6\x7a\x00\xa0\x01\xcc\x11\xdd\x80\x78\xa9\xfd\x64\xb9\xd0\x72\x1c\x5d\xb0\x14\x19\x3e\x34\xe7\x96\x6e\x1a\x96\x0d\x03\xd3\x33\x3c\xe8\x7b\xf5\x31\xff\x08\x23\x93\xc3\xe8\x1e\xf4\x5d\x57\x9c\xf3\xf5\x36\x84\x1f\x99\xd0\x9d\x87\xbc\x16\x09\x5d\x47\x37\xa1\x69\xfb\xba\x05\x5d\xcb\x65\xc7\x01\x00\x60\x6f\xd8\x90\xc1\x19\x26\xb4\xe7\x73\xdd\x84\xd6\xdc\x35\x4c\xe8\x99\x81\x6e\x42\xd7\x9c\x1b\x26\x9c\x5b\x56\x64\x42\xdb\xb3\x0d\x13\x9a\x21\x3b\xef\x39\x16\x3b\xe6\x18\x7d\x9b\xbd\x5b\x6e\x90\xba\xd0\x73\x7c\xdd\x85\x9e\xeb\x73\x78\x06\xd9\x7e\x9f\xfb\x82\x11\xc7\x66\xe5\x99\xc1\xc0\x3d\x06\xee\xf9\x91\x61\x42\x2b\xe4\x78\x82\xc0\x30\xa1\xeb\xb3\x63\xdb\x65\xd4\x7d\xdb\xe3\xf8\x1d\x06\x64\xbb\x8c\x09\xcb\x64\x2c\x0a\xe6\x1c\xf7\xe9\x78\x6e\x86\xff\xf6\xa1\xeb\xb3\x85\xb7\xc1\x67\xce\xb8\x0c\xd8\xa5\xd0\x65\x1c\x5b\xb6\x67\xf0\xd9\x8b\xe3\xc8\x70\x05\x29\x23\x80\x96\xe5\xe8\x0e\x9c\x3b\x61\x75\xcc\xde\xad\xc8\xd4\x19\x84\xc5\xae\xb8\x96\xfe\x74\xa5\x7a\x2f\x39\x94\xc1\xae\x56\x23\xc4\x68\xc6\x00\x93\xa8\x33\xe7\xf3\x08\x3d\xc6\x1c\x13\xee\x7c\x6e\xf3\xc3\x8d\x03\x9d\x79\x18\x31\x08\x8f\x0b\x65\x3e\xb7\x39\xa4\xdf\x1c\xcf\xe7\x37\xf6\x1c\x86\x8e\xaf\xdb\x36\xf4\x2c\x4f\xb7\x2d\x68\x05\x61\x65\x1d\x7a\xc7\x56\xfe\xd0\xc0\x2a\x49\xd3\xa5\xf6\xf2\x07\xfe\x3a\xbb\x7d\xf8\xcf\x1b\xe9\x6c\x08\x52\x0d\x71\xcd\xc2\x80\xec\xfa\x25\x09\x5a\x91\xae\x79\x8b\xf0\x99\xe9\x9e\x23\xe0\x7b\x2f\xca\xca\xa6\x81\xba\xe7\x51\x8d\x64\x51\x9a\x44\x9f\x97\x1a\xaf\x88\x3f\x10\x14\x4f\xaf\x34\x40\x13\xca\x42\xcf\x47\x8a\x0a\x0a\x48\xc1\x6f\xf8\x4e\xb2\x1d\x1e\x1b\x5b\x3b\xed\xd6\xbf\x4b\x74\x75\x6c\x68\x79\x4e\xfd\x39\x26\xbe\x76\x47\x7c\x93\x08\xcb\xf9\xbf\x4b\x91\xb4\xda\xe9\x80\xf3\x38\x5b\xf1\xcd\xdc\xc5\x28\x76\x29\x5e\xe0\x3d\xce\x48\x1c\xbf\x89\xd2\x24\xef\x9e\xd1\x78\x50\xb6\x4d\x07\xfa\xa1\xa5\x87\x21\x9c\xdb\xc1\x07\x37\x84\x8e\xe9\xeb\x36\xb4\x43\x97\xc5\x0f\x6f\xee\x1b\x36\xf4\xad\xd0\xb0\x4c\x68\x3b\x81\x61\xb3\x38\x6a\x58\x2e\xf4\x6c\x4b\x7c\xe1\xe1\xc8\xf2\xa1\xe5\xd8\x2c\x9a\xfb\xd0\xf4\xe6\xba\xe5\x40\xdb\xf6\x9b\x6f\x73\xe8\xf9\xc1\xde\x0a\x03\x18\xba\x91\xa9\xdb\x30\x70\x1c\x16\xbc\x7c\x4f\x5c\x0b\xf5\x36\x64\x18\xb9\xd0\x0e\x1c\xdd\xd4\x43\x18\xf2\x25\xba\xeb\x59\x7a\x8b\x26\x0b\x81\xa9\xe5\xb9\xd0\x0d\x1c\x23\xf4\xa1\xe3\x38\x91\x65\xc3\x79\x60\xf8\xd0\x73\x3d\xdd\x32\xa1\x1b\x84\x86\x35\x87\xae\x1b\x76\xbf\x7d\xb4\xd9\xa7\x6f\xe9\x96\xe9\x43\xb6\xc2\xec\x0a\xe0\xe2\x50\x33\x26\x2e\xc8\xca\xa6\xa1\x15\x01\xaf\xa7\xaa\xad\xb6\x96\x37\xb2\x53\x27\xce\x78\x87\x76\xe5\xb0\x03\xfe\xdd\x3d\x70\xee\x87\xf5\xe7\x48\x0f\x6c\x8d\xf8\x86\x1e\x38\x24\xdc\x06\xfe\x79\x2e\x18\x98\xd0\x63\xb5\xc6\x8f\x8e\x07\xfd\xd0\x8f\x8c\x10\x06\x81\x27\x7c\x29\x08\x03\x3d\x80\xa6\xe5\xd6\x5f\xc4\xc7\xde\x0a\xe7\x30\x08\x44\x26\xd6\x05\x3c\x03\x73\xf4\x0e\x58\xf5\xb1\x71\x5d\xe8\xbb\xf3\xa8\xc2\x5b\x9d\x35\xda\x03\x2a\xf4\xff\x16\x1f\x37\x61\x00\x5d\x57\x10\xd6\x43\x13\xba\x36\xab\x6d\x6a\x3e\xff\x00\x3f\x09\xf2\xb6\x6e\x02\x00\x36\x46\x85\x9e\xf3\x3d\xaf\xf9\x0e\xdb\x7c\x87\x27\x7c\x2b\x98\x0e\xbb\x4c\xb3\x19\x8e\xe7\xdb\xb6\x5c\xe8\x07\x56\xc5\xb9\x6d\xce\xa1\xcf\xcb\xb2\x86\xdf\x6f\xe2\xdf\x6a\xf7\x56\x38\xb8\x48\xb8\xf5\xb6\x78\xcb\xc5\xf9\xb9\x13\x1f\xbf\x41\x59\x84\x53\xb0\xcb\xd9\x62\xe6\xff\x75\xb2\xb5\x6c\x18\x06\x76\xfd\x39\xc6\xd5\xbb\x23\xbe\x9d\xab\xdf\xa4\xe4\x9b\xba\xba\xe5\x58\x30\x30\x5d\xdd\x32\xe7\xd0\x0d\xad\xd4\xf7\x60\xe8\xcc\x0d\xf1\x11\xcd\x61\x18\x1a\xec\x4d\xe7\x47\x56\x00\x1d\xb6\xfa\x30\x6c\x0f\x3a\x96\xcd\xf3\x6c\x03\x22\xae\xda\xe2\x58\x00\xe8\x66\x2a\x30\xf9\xba\xf8\xf8\xe0\x98\xd0\xf3\x5c\xdd\x63\x0b\x8b\xa1\xb1\xfc\x7a\x50\x91\x16\x87\x35\x03\x7a\xc3\x40\xda\xc6\x3f\xff\xc0\x11\xeb\x56\x60\xb3\x08\x31\x8c\x20\x6a\xe6\x26\xae\xd8\xe2\xb8\xe1\x5e\x20\xaf\xe6\xd0\x21\xc5\xd7\x3c\x4f\xe8\x7b\x87\xf7\x8a\xcf\x6e\xc4\xf7\xe1\x44\xf6\x7f\xaf\x48\xc0\x8c\x4f\xb4\x40\x07\x2d\x95\x79\x77\xc7\x9f\x3c\x13\xb8\x9e\x06\x5a\xad\x41\x71\x4b\x8a\xc6\x1f\x7f\x69\x6e\x40\x31\x48\x91\xac\x93\x6c\xa9\xd9\x5e\x7e\x00\xf6\x08\xa7\x68\x48\x8e\x5d\x46\x01\x70\x5d\x3d\xf4\x07\x40\x74\x60\x94\x34\x10\x1d\x97\x1a\x5b\xe2\x69\xa0\x58\x6a\xb6\x39\xcc\xd4\x58\x52\x2d\xe7\x5b\xb4\xb7\x25\xae\xde\x94\xb4\x20\x9f\xf1\xa2\x79\xc4\x0f\x88\x13\x46\x7d\xb7\x6c\x73\x22\x46\xe5\x06\x15\x05\x3a\x2e\x80\x65\xea\xc0\x7b\xa3\x81\xb1\x4b\xcb\xd6\x54\x7b\x67\x1a\x0c\x4f\x54\x52\x16\xfe\x9d\x66\xfa\x57\xad\x9e\x6b\xb8\x61\x1f\xe9\x9c\xeb\xec\x0f\xac\xb5\xca\x4b\xd6\xf5\x22\x58\xc6\xa3\xac\xb0\x56\x2c\xae\xf9\x0a\x77\x70\x6d\xfc\xb4\xb1\xab\x72\x9d\xba\x81\x69\x7b\x1e\xfb\x5b\x58\xb6\xe3\x7a\x73\x60\x00\x3e\xbc\xde\x64\x06\x4d\x37\x1a\x68\xa8\x2c\x25\x66\x3f\x76\xd7\xa0\xb3\x67\xa5\x2c\x0e\xba\x73\x39\xe2\x52\xeb\xc3\x31\x14\x1a\xfe\x83\x4b\x55\x2c\x94\x33\xdd\xc3\x43\x46\x9e\xc5\xc2\xcf\xe4\x79\x1c\x3c\xf7\xda\xc8\x5e\x7d\x19\x15\x49\xde\xd7\x92\x7e\x4c\xb2\x98\x3c\x42\xae\xfe\xef\x11\x45\x60\x09\xbe\xb0\x2a\x62\x01\xb4\x78\x31\xb3\xec\x70\xe6\x07\x81\xf6\xf5\xcd\x69\xef\xbf\x41\x38\x86\x90\xb6\x2b\x31\xf3\xf9\x24\xa2\xda\xd9\xed\x9c\xb3\x57\xe0\xe6\x97\x9f\x3f\x7e\x7a\xf7\xf3\xa7\x8f\x67\x77\x99\xee\x51\x01\x92\x2c\xa1\x09\x4a\x3f\x6d\x30\x67\xcf\x7c\x03\x66\x33\xf0\x69\x83\xeb\x0b\xe2\x16\xc6\x84\x64\x00\x65\xeb\x14\xeb\x20\xc9\x40\x8c\xd7\x05\xc6\x25\x3c\xc7\x46\x19\x9a\xef\x71\xca\x71\xb9\x0d\x2e\xb4\x25\xbb\x8c\x02\x4a\xaa\x1b\x22\xc5\x43\xe7\xbf\xef\x50\x81\x01\x7a\x20\x3b\x0a\xf0\x1e\x17\x47\x60\xcd\xa1\x0f\xb6\x49\x9a\x26\x25\x8e\x48\x16\x97\x03\xe4\x50\xb6\xde\xa5\xa8\xf8\x90\x6c\x13\x0a\x96\x20\x7c\xe2\x7e\x8b\x0e\xc9\x76\xb7\x05\xd9\x6e\xfb\x80\x0b\x40\x56\x35\x96\x5e\x26\x7a\x50\xb7\xb3\x2a\x58\x82\x98\x44\xbb\x2d\xce\x28\x5c\x63\xfa\x5e\x3c\x20\xff\xaf\xe3\x6d\x3c\xed\x66\xdf\xd3\xbd\xf1\x53\x44\xf6\x58\x4c\xb6\x76\xfe\x58\x10\x43\x96\x91\x47\x00\x96\x1d\xa5\x75\x81\x56\xbb\x8c\xef\x45\x81\x98\xbc\xcb\x92\xed\xf4\xaa\xef\xa9\x8b\x64\x35\xcd\xc8\xe3\xdb\xa5\x33\x37\xaf\x24\x9b\xdd\x8c\xd0\x12\x98\x23\xf7\xcf\xdb\xac\x43\x9e\xba\x7e\x9d\x74\x52\xe1\xe4\x37\xb0\x04\x5a\x75\x37\xac\x06\x5e\x73\xfc\xaf\x81\x16\xe3\xf5\x55\x5f\x18\xe4\x97\x97\x2d\x73\xea\xbb\x79\x16\x94\x29\xc6\xf9\x94\x25\x37\x31\xdb\x93\x5b\x06\xce\x58\xad\x65\x72\x7a\xbe\x11\x1a\x47\x08\xa6\x34\xd9\x62\x1d\x14\xb8\x24\xe9\x1e\xf7\x4a\x90\xbd\xa6\xcd\xb0\xe9\x10\x28\x67\x16\xd3\x4f\xc9\x16\x93\x1d\x6d\x0d\xbc\x02\x5f\x06\x12\x6e\x85\x7a\x2a\xbb\x61\x85\x4f\x54\x67\x1c\xf7\xdc\x30\xc1\x2f\x5e\x35\xec\xf5\x68\x93\xc5\xb8\x3e\x43\x2b\xf0\xef\x3b\x5c\xd2\x77\xfc\x59\x9e\x84\x64\x3f\x14\x68\x8b\x6f\xbf\x07\x12\xc1\xf1\xce\xb0\xd0\x7f\x33\x64\xda\x67\x5b\xb3\x99\x14\xf5\xb2\xff\xca\x74\x9c\x6e\x5b\xac\x88\x56\xd9\x08\x4e\x22\xbe\xbc\x3e\x21\x27\x61\x6f\x88\x7e\xc7\x05\xaa\xc7\xc0\xeb\xa0\xda\x76\xd7\x81\xf8\xca\xc2\x13\x8b\xe7\x8f\x1b\x9c\xf1\x00\x85\x6a\x3e\x40\xc9\x64\x5c\xea\xa0\xa4\xa4\xc0\x31\x0b\x8c\x08\x44\xbb\x92\x92\x2d\xc8\x0b\x92\xe3\x82\x1e\x61\x87\xab\xa7\xc3\x56\x46\x69\xf3\x7c\xfd\xc2\x30\x80\x60\xb9\x04\x02\x02\x18\xc6\xe8\xec\xc6\x23\x70\x8a\x0b\xfa\x3e\x55\xc5\x35\xb1\x89\xd1\x63\x7c\x6c\x3c\xdf\x7a\x50\x8f\x7f\xda\x9d\x90\xe0\x28\x29\xc9\xd5\x28\x9a\x8e\xaa\x8c\x8b\x02\xa3\xcf\x03\x5c\x34\x2d\x1b\x09\x0e\xb1\x94\x53\x23\xa9\x96\x7b\x57\xe7\x3e\x07\xda\x26\xdc\xda\x7b\x91\x04\x14\x9e\xb9\x73\xe5\x94\x9b\x4d\xf9\x2b\xc8\xb7\xdc\x25\xf1\x23\x59\x4d\x4f\xab\x14\xc8\x4a\x94\xe5\x52\xcb\x48\x86\x35\x19\x07\xec\x55\x29\x1f\x26\x59\x86\x8b\x1f\x3f\xfd\xf4\x61\xa9\x89\x9b\x1c\x9e\x0a\xdd\x17\x2f\x5e\x48\xca\xdc\xaf\x38\x2d\x31\x23\xff\x22\xc9\x9f\x47\xe4\xf6\x6e\x00\xfb\x65\x58\x15\xab\xd0\xca\x4a\x45\x6a\x83\xd5\x8d\x60\x60\x09\x26\x4c\x44\x13\xf9\xb8\xfe\xb0\x28\x87\xaf\x6d\xe8\x8c\x90\xd6\xbe\xe7\x4c\xc1\xa8\xfc\x4a\xa5\x64\x7c\xa0\xb8\xc8\x10\x9b\xf8\x9e\x7c\xc6\xd3\x49\x59\xb3\xb8\xf8\xa2\x31\xbd\x6b\x0b\xa0\x4d\x5e\xf7\x9a\xc4\xeb\x89\xa6\x03\x2d\xc9\x05\x48\x92\x83\xd7\x13\xed\xeb\x44\x96\x76\x46\x96\x0e\x4f\xb9\xb7\xd9\xe2\x90\xa8\x4d\xb8\xf9\x45\x3a\x90\xea\xad\x2d\xce\xc9\xe8\xe7\xaf\x1b\x5e\x5b\xbd\xda\xbf\x90\x59\x11\x84\xbe\xc9\x04\x2f\x9b\x1e\x8f\x3f\xef\xa2\x08\xe7\x14\xcb\xa7\xd8\x9b\x68\xfb\x41\xe5\x76\xcd\x63\xcc\x65\x22\x1c\x98\x9a\x52\x94\x7f\x5e\x2c\xa2\x25\x8f\xe3\x69\x81\x51\x49\x32\x99\x70\xce\xe2\x0b\x2f\x6f\xd8\x88\xbf\x50\x42\x7f\x99\xf6\x55\xf5\xc1\xcf\xf8\x11\x7c\xaa\xfa\x17\xbf\x30\x09\x94\xf2\x1a\x01\x00\x51\x4f\x9e\x12\x64\xe9\x4a\x74\x42\x14\x19\x4b\xf4\x63\x24\xd9\xf5\xa9\x0d\x33\x88\x42\x34\x6b\x94\x78\x8e\xb8\x1c\x46\x73\xc4\xa5\x1a\x4b\x46\x86\x91\x64\xa4\x0f\x47\x9f\x8c\x1a\xa2\x10\xc5\xf1\xfb\x3d\xce\xe8\x87\xa4\xa4\x38\xc3\xc5\x74\xc2\xb7\x88\x26\x7a\x63\x8a\x52\x9f\x94\x45\x79\xc4\x9d\xb9\x52\x23\xd7\xe2\x62\xf2\x9a\x53\x2c\x3f\xb2\x7a\x37\x8b\xf0\xaf\xdd\xaf\x30\xc5\xd9\x9a\x6e\x80\x01\xac\xdf\xe0\xed\x9d\x6c\xf9\x91\xe1\x03\xed\x5d\x9b\x7c\xed\x9d\x76\xff\xa4\x33\xf2\x2d\xe6\x2c\xaa\xfb\xff\xeb\x39\x2b\x0c\xa6\xa1\x03\x96\xe0\xd7\xdf\x24\xb5\xec\xee\x81\xf9\xd2\x03\x73\x33\x01\xa5\x4a\x99\x35\xf0\x27\x22\x3c\x93\xaf\x2c\xa5\x29\xe9\x09\x35\xcc\x77\xe5\x46\x00\x5f\x1a\x04\x77\x59\x83\x47\x49\x6d\x36\xeb\x12\x24\xf9\xb4\xfd\x9d\xff\x96\xcb\x2f\x2b\x81\x42\x92\x35\xda\xf0\x65\x9e\x26\x11\x96\xa3\xd0\x81\x25\x79\x30\x61\x28\xa0\x97\x77\x4c\x16\xfc\x90\xcd\xa5\x9f\x95\x13\x3b\xc9\x9f\x86\x48\xab\xea\x7e\xd3\x5a\x2e\x81\xa5\xaa\x76\x65\xc6\x1c\x27\x28\x25\xeb\x24\x5b\x91\x05\x0b\xc8\x75\x43\x99\x13\x91\x95\x60\xa7\x32\x5c\x91\xe2\x3d\x8a\x36\x4d\xfb\x41\xa9\xbf\xb6\xb0\x4e\xe6\xf2\xab\xf9\x9b\x82\x62\x9f\xe3\x03\xe9\xaf\x50\x5e\xaa\x32\xe1\x7b\xb2\x07\xf8\x18\x9b\x55\x36\xa4\x84\x2d\x8f\x26\xfc\xd1\xaf\xde\x0a\x0f\x28\xd4\xb4\x5c\x9a\x57\x5f\x0a\x4c\x77\x45\x26\xf9\xf1\xcc\x93\x61\x95\x79\x9a\x12\x33\x54\xd1\x02\x6f\x81\xb4\x0b\x07\xfe\xee\x1a\xec\x55\x56\x7f\x58\xe2\xfc\xa6\xa8\x14\xa1\x58\xa9\x42\x9e\xc1\x3b\xa5\x53\x33\x0e\x52\x7c\xe8\xf9\x29\x23\xd0\xaf\x7d\x27\x10\xda\xef\x03\x1f\x13\xbe\x65\x75\x91\xac\x25\x22\xd6\xf5\x7d\x32\x91\x95\x09\xd5\x83\x0a\x93\xab\xf3\x0c\xa8\xf1\x0c\xa8\xe9\x40\x80\xfc\xc0\x54\xdd\x63\x5b\xd2\x02\xa4\x75\x0b\xbc\x0a\xbb\x00\xbb\x21\x59\x86\x23\x8a\xe3\xdb\xbb\xb2\x47\xfd\xb3\x99\x7c\x06\x4d\x7f\xa6\x6f\x12\x4d\x1a\xe7\xeb\xcc\x51\x35\x41\xfd\x8c\xca\x34\x30\x65\x9c\xd4\x6d\xdc\x21\x54\xb2\x68\xba\x3d\xde\xde\x4d\x7a\x93\xb5\xa4\xb7\x83\x4a\x7a\xbf\x4a\x32\xbe\x41\x32\x66\x50\x13\xaf\xce\xa5\x7b\x71\x1d\xb3\xc6\xb4\x3d\x7e\x72\xd5\x4f\x52\x1d\x37\x0b\x1c\x17\xe8\x91\x91\x4f\xf2\x52\xc6\x41\x8a\x29\x10\x7d\x23\xa2\xaa\xd4\x27\x4d\x6f\x49\x96\x79\x6a\x24\x1d\x1f\x96\x35\x58\x18\xd5\x0c\x3f\xbe\x4f\x65\x81\x33\xef\x09\x7c\xd5\x8f\x14\xab\x62\x1f\xc7\xd9\x9e\x46\x54\x60\x44\x71\x35\x93\xe9\x44\x3c\x67\xa2\x4a\x9e\x1c\x83\xe8\x9e\x81\x65\xfd\xbb\xc8\x43\xe0\x7c\xca\x9f\xf0\x81\x3e\x0d\xe1\xe9\x81\x4e\x27\x8b\xc9\xd5\xaf\x66\x4f\xe1\x77\x26\x36\x94\xe7\x38\x8b\x6f\x36\x49\x1a\x4f\x39\x56\x69\x6f\x7f\x5c\xcd\x73\xf2\xbd\xc7\x52\xc6\xf9\x42\x63\x4e\x8d\x8f\xee\x51\xaa\xea\x51\x0a\x9f\x71\xe0\xfc\x15\xd8\xa3\x54\x9a\x18\x57\x49\xf6\xd6\x0a\x94\x69\x70\x60\xdb\x4c\x3c\x71\x77\x55\x87\xff\xe6\x07\x93\x9f\x76\x9d\xac\xc0\x94\xed\x36\x8d\x24\x21\x1e\xbe\x53\xd1\xd0\x5e\xaf\x92\xec\xb5\x74\x57\x6b\x88\x4e\xfd\x24\x62\x4d\x22\x4a\x93\x7c\xa9\xf1\x9f\x8f\x46\x3b\x4a\xf4\xee\xdb\x95\xea\x6e\x92\x56\xc4\x5a\x25\xb2\x5f\x0b\x06\xd2\x8e\xe1\x50\xfb\xf4\x39\x53\x30\xf3\x83\x2e\x7e\xfa\x5a\xbc\xf3\x9f\x35\xff\x13\xea\x18\xd4\xf8\x9f\xd4\xc6\x68\xad\x0f\xd9\xd5\x28\x5d\x48\xea\x4c\x69\x0c\x3e\x7d\x92\x73\x72\xd5\x09\xb6\x7b\x94\xbe\xd6\xbe\xeb\x61\x6a\xa0\x09\x8b\xe9\x4f\xc7\xdb\xbb\xa9\xbc\x25\x2f\x15\x17\x7f\x06\xb8\xcb\x44\x92\x3f\x83\xfe\x07\xb2\x9e\xa6\x64\x7d\x31\xfd\x94\xac\x4f\xc8\xa7\x64\xdd\x4b\xff\xf4\x94\x82\x9f\xa7\xd2\xeb\xe2\xb4\x4d\x58\x08\x4f\x8a\x89\x62\x3d\xd0\xdb\x65\x6f\x2f\x41\x06\xd7\x89\xdd\xa1\x03\xed\x42\x89\x89\xb1\x38\x5d\x8d\xee\x47\xfb\x2c\xfe\xdf\x2e\x2d\xf3\xea\x4b\x85\x77\x02\x21\x9c\xbc\xe6\x57\xd9\xe2\x80\x16\xd3\x16\x24\x30\x80\x7f\x75\xa9\xf9\x9f\x3f\xdf\xdb\xd5\x3d\xc3\x7f\x71\x3f\x80\x57\xb2\x4a\x65\x4b\x30\x2a\x57\x0d\xd7\xb3\xea\x19\xfb\xeb\x19\xff\xaf\x36\xfe\x37\x00\x00\xff\xff\x01\x62\xd9\x23\x82\x63\x00\x00")

func dataIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_dataIndexHtml,
		"data/index.html",
	)
}

func dataIndexHtml() (*asset, error) {
	bytes, err := dataIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/index.html", size: 25474, mode: os.FileMode(438), modTime: time.Unix(1527957220, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"data/index.html": dataIndexHtml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"data": &bintree{nil, map[string]*bintree{
		"index.html": &bintree{dataIndexHtml, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

