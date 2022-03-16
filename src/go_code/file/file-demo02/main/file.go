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
	file, err := os.Open(fileOperation.path)
	if err != nil {
		fmt.Println("Open File Error =", err.Error())
	}
	fileOperation.file = file
}

func (fileOperation *FileOperation) readFile() {
	if fileOperation.file == nil {
		return
	}
	reader := bufio.NewReader(fileOperation.file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(str)
	}
}

func (fileOperation *FileOperation) CloseFile() {
	if fileOperation.file != nil {
		err := fileOperation.file.Close()
		if err != nil {
			fmt.Println("Close file err =", err.Error())
		}
	}
}

func main() {
	fo := FileOperation{path: "/home/ezrachai/Documents/Go.txt"}
	fo.OpenFile()

	fo.readFile()

	defer fo.CloseFile() //	Must CLOSE the file to avoid memory leakage
}
