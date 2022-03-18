package model

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Person struct {
	Name    string
	Age     int
	Address string
}

func NewPerson(personChan chan *Person) {
	go rand.Seed(time.Now().Unix())
	for i := 1; i <= 10; i++ {
		personChan <- &Person{
			Name:    "person" + strconv.Itoa(i),
			Age:     rand.Intn(100) + 1,
			Address: "random address",
		}
	}
	close(personChan)
}

func ReadPerson(personChan chan *Person, exitChan chan bool) {
	for person := range personChan {
		fmt.Println(*person)
	}
	exitChan <- true
	close(exitChan)
}
