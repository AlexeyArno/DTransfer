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

var dirsAmount uint64
var filesAmount uint64

// history of crawler
var visited []string

// Clear history
func Clear() {
	visited = visited[:0]
}

//SetPath - set new path to dir
func SetPath(p string) {
	currentPath = ""
	beginPath = p
	dirsAmount = 0
}

// pathBack - return to the previous level in the file system hierarchy
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
		if currentPath[len(currentPath)-1-i] == '/' {
			currentPath = currentPath[:len(currentPath)-1-i]
			return
		}
	}
}

// GetStat return amount Files and Dirs
func GetStat() (uint64, uint64) {
	return filesAmount, dirsAmount
}

// inVisited - return is there a file in the history
func inVisited(path string) bool {
	for _, j := range visited {
		if path == j {
			return true
		}
	}
	return false
}

// Next - simple generator for crawling directories
func Next() (string, bool, error) {
	files, err := ioutil.ReadDir(beginPath + "/" + currentPath)
	if err != nil {
		return "", false, err
	}
	visitedCount := 0

	for _, f := range files {
		if !inVisited(currentPath + "/" + f.Name()) {
			visited = append(visited, currentPath+"/"+f.Name())
			if f.IsDir() {
				dirsAmount++
				currentPath = currentPath + "/" + f.Name()
				return currentPath, true, nil
			} else {
				filesAmount++
				return currentPath + "/" + f.Name(), false, nil
			}
		} else {
			visitedCount++
		}
	}
	if visitedCount == len(files) {
		if currentPath == "" {
			Clear()
			return "", false, errors.New("All visited")
		}
		pathBack()
		return Next()
	}
	Clear()
	return "", false, errors.New("All visited")
}
