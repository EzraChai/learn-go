package main

import (
	"flag"
	"fmt"
)

type arguments struct {
	length   int
	user     string
	password string
	host     string
	port     int
}

func NewArguments() *arguments {
	return &arguments{}
}

//	Useful ✅
func main() {
	args := NewArguments()
	flag.StringVar(&args.user, "u", "", "Username")
	flag.StringVar(&args.password, "pwd", "", "Password")
	flag.StringVar(&args.host, "h", "localhost", "Host")
	flag.IntVar(&args.port, "p", 8080, "Port")

	//	Important Function, Parse
	flag.Parse()

	// RUN FILE : ./args -u admin -pwd 123456 -h localhost -p 8000
	//OUTPUT : user=admin password=123456 host=localhost port=8000

	fmt.Printf("user=%v password=%v host=%v port=%v", args.user, args.password, args.host, args.port)
}
