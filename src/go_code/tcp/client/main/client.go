package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	dial, err := net.Dial("tcp", "192.168.0.109:8888")
	if err != nil {
		panic(err)
	}
	defer func(dial net.Conn) {
		err := dial.Close()
		if err != nil {
			panic(err)
		}
	}(dial)
	for {
		fmt.Printf("word: ")
		reader := bufio.NewReader(os.Stdin)
		line, err := reader.ReadString('\n')
		if line == "exit\n" {
			break
		}
		if err != nil {
			panic(err)
		}
		_, err = dial.Write([]byte(line))
		if err != nil {
			return
		}
	}
	fmt.Println(dial.LocalAddr())
}
