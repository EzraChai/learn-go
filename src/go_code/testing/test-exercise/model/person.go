package model

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Person struct {
	Name  string
	Age   int
	Hobby string
}

func (person *Person) SerializePerson() (err error) {
	marshal, err := json.Marshal(person)
	if err != nil {
		return
	}
	file, err := OpenFile("./persistence.txt")
	if err != nil {
		return
	}
	defer func(file *os.File) {
		err = CloseFile(file)
		if err != nil {
			return
		}
	}(file)

	WriteFile(file, marshal)
	return
}

func (person *Person) DeserializePerson() (err error) {
	file, err := OpenFile("./persistence.txt")
	defer func(file *os.File) {
		err = CloseFile(file)
		if err != nil {
			return
		}
	}(file)
	jsonBytes := ReadFile(file)
	err = json.Unmarshal(jsonBytes, person)
	if err != nil {
		return
	}
	return
}

func ReadFile(file *os.File) []byte {
	reader := bufio.NewReader(file)
	for {
		readBytes, err := reader.ReadBytes('\n')
		if err == io.EOF {
			return readBytes
		}
	}
}

func WriteFile(file *os.File, json []byte) {
	writer := bufio.NewWriter(file)
	for _, b := range json {
		err := writer.WriteByte(b)
		if err != nil {
			fmt.Println(err)
		}
	}
	err := writer.Flush()
	if err != nil {
		return
	}
}

func CloseFile(file *os.File) (err error) {
	err = file.Close()
	if err != nil {
		return
	}
	return
}

func OpenFile(path string) (file *os.File, err error) {
	file, err = os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return
	}
	return
}
