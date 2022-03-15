package main

import "fmt"

type Monkey struct {
	Name string
}

type BirdAble interface {
	Fly()
}

type FishAble interface {
	Swim()
}

func (monkey *Monkey) Climbing() {
	fmt.Println(monkey.Name + ": \"I'm climbing.\"")
}

type LittleMonkey struct {
	Monkey
}

func (monkey *LittleMonkey) Fly() {
	fmt.Println(monkey.Name + ": \"I'm flying.\"")
}

func (monkey *LittleMonkey) Swim() {
	fmt.Println(monkey.Name + ": \"I'm swimming.\"")
}

func main() {
	monkey := LittleMonkey{
		Monkey{
			Name: "MONEY",
		},
	}
	monkey.Climbing()
	monkey.Fly()
	monkey.Swim()
}
