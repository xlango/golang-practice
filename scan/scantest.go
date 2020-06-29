package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

var count = 0

func logFileName(path string, info os.FileInfo, err error) error {

	count++

	// 返回错误后，编辑将终止
	if count > 10 {
		return errors.New("stop")
	}

	fmt.Printf("\n %d path: %s    fileName: %s", count, path, info.Name())
	return nil

}

func main() {
	filepath.Walk("/", logFileName)
}
