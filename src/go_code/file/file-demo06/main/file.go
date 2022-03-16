package main

import (
	"bufio"
	"fmt"
	"os"
)

type FileOperation struct {
	path string
	file *os.File
}

func (fileOperation *FileOperation) OpenFile() {
	file, err := os.OpenFile(fileOperation.path, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("open file error :", err.Error())
	}
	fileOperation.file = file
}

func (fileOperation *FileOperation) CloseFile() {
	err := fileOperation.file.Close()
	if err != nil {
		return
	}
}

func (fileOperation *FileOperation) WriteFile() {
	str := "ABC! ENGLISH!\r\n"
	writer := bufio.NewWriter(fileOperation.file)
	for i := 0; i < 5; i++ {
		_, err := writer.WriteString(str)
		if err != nil {
			return
		}
	}
	err := writer.Flush()
	if err != nil {
		return
	}
}

func main() {
	fo := FileOperation{path: "/home/ezrachai/GolandProjects/learn-go/src/go_code/file/go.txt"}
	fo.OpenFile()

	fo.WriteFile()

	defer fo.CloseFile()
}
