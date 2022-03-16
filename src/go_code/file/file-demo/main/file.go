package main

import (
	"fmt"
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
	fmt.Println(fileOperation.file)
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
	fo := FileOperation{path: "/home/ezrachai/Pictures/Screenshot from 2021-10-11 20-41-13.png"}
	fo.OpenFile()
	defer fo.CloseFile()
}
