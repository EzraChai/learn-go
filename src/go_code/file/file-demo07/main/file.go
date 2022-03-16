package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type FileOperation struct {
	path string
	file *os.File
}

func (fileOperation *FileOperation) OpenFile() {
	file, err := os.OpenFile(fileOperation.path, os.O_RDWR|os.O_APPEND, 0666)
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
	str := "I'm from MALAYSIA!\r\n"
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

func (fileOperation *FileOperation) ReadFile() {
	reader := bufio.NewReader(fileOperation.file)
	for {
		str, err := reader.ReadString('\n')
		//	End Of File
		if err == io.EOF {
			return
		}
		fmt.Print(str)
	}
}

func main() {
	fo := FileOperation{path: "/home/ezrachai/GolandProjects/learn-go/src/go_code/file/go.txt"}
	fo.OpenFile()
	fo.ReadFile()
	fo.WriteFile()

	defer fo.CloseFile()
}
