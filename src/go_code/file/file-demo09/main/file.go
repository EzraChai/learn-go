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

type CharCount struct {
	CharacterCount int
	NumberCount    int
	SpaceCount     int
	OtherCount     int
}

func (fileOperation *FileOperation) OpenFile() {
	file, err := os.Open(fileOperation.path)
	if err != nil {
		panic(err)
		return
	}
	fileOperation.file = file
}

func (fileOperation *FileOperation) CloseFile() {
	err := fileOperation.file.Close()
	if err != nil {
		return
	}
}

func (fileOperation *FileOperation) ReadFile() *CharCount {
	var count = &CharCount{}
	reader := bufio.NewReader(fileOperation.file)
	for {
		str, err := reader.ReadString('\n')

		//	Iterate str
		for _, v := range []rune(str) {
			fmt.Print(string(v))
			switch {
			case v >= 'a' && v <= 'z':
				count.CharacterCount++
			case v >= 'A' && v <= 'Z':
				count.CharacterCount++
			case v == ' ' || v == '\t':
				count.SpaceCount++
			case v >= '0' && v <= '9':
				count.NumberCount++
			default:
				count.OtherCount++
			}
		}
		if err == io.EOF {
			break
		}
	}
	return count
}

func main() {
	fo := FileOperation{path: "/home/ezrachai/GolandProjects/learn-go/src/go_code/file/file-demo09/go.txt"}
	fo.OpenFile()
	count := fo.ReadFile()
	fmt.Println(*count)
	defer fo.CloseFile()
}
