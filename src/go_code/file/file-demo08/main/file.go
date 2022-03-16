package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

type FileOperation struct {
	OriginPath      string
	DestinationPath string
	file            []byte
}

func (fileOperation *FileOperation) ReadFile() {
	stat, err := os.Stat(fileOperation.OriginPath)
	if err != nil {
		fmt.Println(err)
	}
	var str string = fileOperation.FormatStat(stat)
	fmt.Println(str)
	file, err := ioutil.ReadFile(fileOperation.OriginPath)
	if err != nil {
		return
	}
	fileOperation.file = file
}

func (fileOperation *FileOperation) WriteFile() {
	_, err := os.Stat(fileOperation.DestinationPath)
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile(fileOperation.DestinationPath, fileOperation.file, 0666)
	if err != nil {
		return
	}
}

func (fileOperation *FileOperation) FormatStat(stat os.FileInfo) string {
	return fmt.Sprintf("Filename: [%v] FileSize: [%v] FileMode: [%v],", stat.Name(), stat.Size(), stat.Mode().String())
}

func main() {
	fo := FileOperation{
		OriginPath:      "/home/ezrachai/GolandProjects/learn-go/src/go_code/file/file-demo08/origin.txt",
		DestinationPath: "/home/ezrachai/GolandProjects/learn-go/src/go_code/file/file-demo08/dest.txt",
	}
	fo.ReadFile()
	fo.WriteFile()
}
