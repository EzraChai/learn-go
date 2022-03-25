package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("Server is listening")
	//	1.	tcp: used tcp as the network protocols
	//	2.	0.0.0.0:8888: listening in local computer port 8888
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		panic(err)
	}
	defer func(listen net.Listener) {
		err := listen.Close()
		if err != nil {
			panic(err)
		}
	}(listen)
	for {
		//	Waiting for the client to be connected
		fmt.Println("Waiting...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept():", err)
		} else {
			fmt.Println("Remote address:", conn.RemoteAddr().String())

			//	goroutine here to serve the client
			go process(conn)
		}
	}
}

func process(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(conn)

	for {
		//	Create a new slice
		buf := make([]byte, 1024)
		//	conn.Read(buf)
		//	1.	Waiting for the client to send message through conn
		//	2. 	If client doesn't write, goroutine will block here
		//fmt.Println("Waiting... ", conn.RemoteAddr().String())
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Printf("%e", err)
			return
		}
		fmt.Print(string(buf[:n]))
	}
}
