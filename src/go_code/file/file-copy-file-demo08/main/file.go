package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type FileOperation struct {
	OriginPath string
	DestPath   string
}

func CopyFile(destFileName, srcFileName string) (written int64, err error) {
	file, err := os.OpenFile(srcFileName, os.O_RDWR, 0666)
	err1 := file.Close()
	if err1 != nil {
		return -1, err1
	}
	if err != nil {
		return -1, err
	}
	reader := bufio.NewReader(file)
	destFile, err := os.OpenFile(destFileName, os.O_WRONLY|os.O_CREATE, 0666)
	err2 := destFile.Close()
	if err2 != nil {
		return -1, err2
	}
	if err != nil {
		return -1, err
	}
	writer := bufio.NewWriter(destFile)

	return io.Copy(writer, reader)
}

func main() {
	fo := &FileOperation{
		OriginPath: "/home/ezrachai/GolandProjects/learn-go/src/go_code/file/file-copy-file-demo08/chloe-gan.jpg",
		DestPath:   "/home/ezrachai/GolandProjects/learn-go/src/go_code/file/chloe-gan.jpg",
	}
	_, err := CopyFile(fo.DestPath, fo.OriginPath)
	if err != nil {
		fmt.Println(err)
	}
}
