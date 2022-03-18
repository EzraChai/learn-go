package model

import (
	"fmt"
	"testing"
)

func TestNewPerson(t *testing.T) {

	personChan := make(chan *Person, 10)
	exitChan := make(chan bool, 1)

	go NewPerson(personChan)

	go ReadPerson(personChan, exitChan)
	for {
		v := <-exitChan
		if v {
			break
		}
	}
	fmt.Println("Haloo Chloe")
}
