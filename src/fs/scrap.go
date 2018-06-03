package fs

import (
	"errors"
	"io/ioutil"
	"strings"
)

var currentPath string
var lastFile string
var beginPath string

var limiter uint64 = 100

var counter uint64

var Dirs uint64 = 0
var Files uint64 = 0

var visited []string

func SetPath(p string) {
	currentPath = ""
	beginPath = p
	Dirs = 0
}

func pathBack() {
	deleteIndexes := []int{}
	for i, v := range visited {
		if strings.HasPrefix(v, currentPath) {
			deleteIndexes = append(deleteIndexes, i)
		}
	}
	for i := len(deleteIndexes) - 1; i >= 0; i-- {
		visited = visited[:deleteIndexes[i]+copy(visited[deleteIndexes[i]:], visited[deleteIndexes[i]+1:])]
	}

	visited = append(visited, currentPath)

	for i := range currentPath {
		if currentPath[len(currentPath)-1-i] == '\\' {
			currentPath = currentPath[:len(currentPath)-1-i]
			return
		}
	}
}

func GetStat() (uint64, uint64) {
	return Files, Dirs
}

func inVisited(path string) bool {
	for _, j := range visited {
		if path == j {
			return true
		}
	}
	return false
}

func Next() (string, bool, error) {
	if counter > 100 {
		counter = 0
		return "", false, errors.New("Limit")
	}
	files, err := ioutil.ReadDir(beginPath + "\\" + currentPath)
	if err != nil {
		return "", false, err
	}
	visitedCount := 0
	for _, f := range files {
		if !inVisited(currentPath + "\\" + f.Name()) {
			visited = append(visited, currentPath+"\\"+f.Name())
			if f.IsDir() {
				Dirs++
				currentPath = currentPath + "\\" + f.Name()
				return beginPath + currentPath, true, nil
			} else {
				Files++
				return beginPath + currentPath + "\\" + f.Name(), false, nil
			}

		} else {
			visitedCount++
		}
	}
	if visitedCount == len(files) {
		if currentPath == "" {
			return "", false, errors.New("All visited")
		}
		pathBack()
		return Next()
	}
	counter++
	return "", false, errors.New("All visited")
}
