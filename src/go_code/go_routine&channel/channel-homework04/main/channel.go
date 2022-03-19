package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type FileOperation struct {
	path string
	file *os.File
}

func (fileOperation *FileOperation) OpenFile() {
	file, _ := os.OpenFile(fileOperation.path, os.O_RDWR, 0666)
	fileOperation.file = file
}

func (fileOperation *FileOperation) CloseFile() {
	err := fileOperation.file.Close()
	if err != nil {
		fmt.Println(err)
	}
	return
}

func (fileOperation *FileOperation) writeDataToFile(chanBool chan bool) {
	writer := bufio.NewWriter(fileOperation.file)
	for i := 1; i < 1000; i++ {
		_, err := writer.WriteString(getRandomNumber())
		if err != nil {
			fmt.Println(err)
		}
	}
	err := writer.Flush()
	chanBool <- true
	close(chanBool)

	if err != nil {
		fmt.Println(err)
	}
}

func (fileOperation *FileOperation) sort(chanBool chan bool) {
	reader := bufio.NewReader(fileOperation.file)
	for {
		readString, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		readString = strings.Trim(readString, "\n")
		fmt.Println("Str", readString)
		_, err = strconv.Atoi(readString)
		if err != nil {
			fmt.Println(err)
			break
		}
	}
	chanBool <- true
	close(chanBool)
}

func getRandomNumber() string {
	return strconv.Itoa(rand.Intn(1000)) + "\n"
}

func main() {
	chanBool := make(chan bool, 1)
	chanBool2 := make(chan bool, 1)
	go rand.Seed(time.Now().Unix())
	fo := &FileOperation{path: "./persistence.txt"}
	fo.OpenFile()
	go fo.writeDataToFile(chanBool)
	for {
		value := <-chanBool
		if value {
			break
		}
	}
	fo.CloseFile()
	fo2 := &FileOperation{path: "./persistence.txt"}

	go fo2.sort(chanBool2)
	for {
		value := <-chanBool2
		if value {
			break
		}
	}
	fo2.CloseFile()

}
