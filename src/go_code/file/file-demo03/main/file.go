package main

import (
	"fmt"
	"io/ioutil"
)

type FileOperation struct {
	path string
}

func (fileOperation *FileOperation) readFile() {
	content, err := ioutil.ReadFile(fileOperation.path)
	if err != nil {
		fmt.Println(err)
	}
	str := fmt.Sprintf("%v", string(content)) //	content = []byte
	fmt.Println(str)
}

func main() {
	fo := FileOperation{path: "/home/ezrachai/Documents/Go.txt"}
	fo.readFile()
}
