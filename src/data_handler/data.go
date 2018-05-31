package data_handler

var (
	upload          bool
	download        bool
	currentFilePath string
)

type command struct {
	Data string `json: data`
}

type newDirectory struct {
	Name string `json: name`
	Path string `json: path`
}

type newFile struct {
	Name string `json: name`
	Path string `json: path`
}
