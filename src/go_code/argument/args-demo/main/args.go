package main

import (
	"fmt"
	"os"
)

type arguments struct {
	length   int
	argument []string
}

func (a *arguments) getArgumentDetails() {
	for i, s := range a.argument {
		fmt.Printf("args[%v] = %s\n", i, s)
	}
}

func NewArguments() *arguments {
	return &arguments{
		length:   len(os.Args),
		argument: os.Args,
	}
}

func main() {
	args := NewArguments()
	fmt.Println("Length of args :", args.length)
	args.getArgumentDetails()
}
