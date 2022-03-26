package main

import "fmt"

type PersonNode struct {
	no   int
	name string
	next *PersonNode
}

func InitHeaderNode() *PersonNode {
	return &PersonNode{}
}

//	List out all the PersonNode
func ListPersonNode(head *PersonNode) {
	temp := head
	for {
		fmt.Println(temp)
		if temp.next == nil {
			break
		}
		temp = temp.next
	}
}

//	Add node to the list
func InsertPersonNode(head *PersonNode, newNode *PersonNode) {

	//	Find the last node
	temp := head
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}
	temp.next = newNode
}

func main() {
	head := InitHeaderNode()

	//Create a new Node
	person1 := &PersonNode{
		no:   1,
		name: "Chloe Gan",
	}

	person2 := &PersonNode{
		no:   2,
		name: "Ezra Chai",
	}

	person3 := &PersonNode{
		no:   3,
		name: "Jacky",
	}
	InsertPersonNode(head, person1)
	InsertPersonNode(head, person2)
	InsertPersonNode(head, person3)

	ListPersonNode(head)
}
