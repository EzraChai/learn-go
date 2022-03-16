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
	file, err := os.OpenFile(fileOperation.path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Open file error:", err)
		return
	}
	fileOperation.file = file
}

func (fileOperation *FileOperation) WriteFile() {
	//	Ready to write "Hello, Gordan" 5 times
	str := "Hello, Gordan\r\n"

	//	When writing, use *Writer
	writer := bufio.NewWriter(fileOperation.file)

	for s := 0; s < 5; s++ {
		_, err := writer.WriteString(str)
		if err != nil {
			return
		}
	}

	//	Content is written into a memory, so you need to use Flush() to write the file
	err := writer.Flush()
	if err != nil {
		return
	}
}

func (fileOperation *FileOperation) CloseFile() {
	err := fileOperation.file.Close()
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
