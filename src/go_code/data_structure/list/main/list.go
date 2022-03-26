package main

import (
	"errors"
	"fmt"
)

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
	if temp.next == nil {
		return
	}
	for {
		fmt.Println(temp.next)
		temp = temp.next

		//	Last node
		if temp.next == nil {
			break
		}
	}
}

//	Sort by the no from smallest to biggest
func InsertHeroNodeV2(head *PersonNode, newNode *PersonNode) {
	temp := head
	for {
		if temp.next == nil {
			break
		} else if temp.next.no > newNode.no {
			break
		} else if temp.next.no == newNode.no {
			panic(errors.New("duplicated Id"))
		}
		temp = temp.next
	}
	newNode.next = temp.next
	temp.next = newNode
}

//	Delete a node element
func DeleteNodeElement(head *PersonNode, id int) (err error) {
	temp := head
	for {
		if temp.next == nil {
			err = errors.New("node not found")
			break
		} else if temp.next.no == id {
			temp.next = temp.next.next
			break
		}
		temp = temp.next
	}
	return
}

func EditNodeElement(head *PersonNode, newPerson *PersonNode) (err error) {
	temp := head

	for {
		if temp.next == nil {
			err = errors.New("node not found")
			break
		} else if temp.next.no == newPerson.no {
			newPerson.next = temp.next.next
			temp.next = newPerson
			break
		}
		temp = temp.next
	}
	return
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

	person4 := &PersonNode{
		no:   4,
		name: "Brian",
	}

	InsertHeroNodeV2(head, person4)
	InsertHeroNodeV2(head, person3)
	InsertHeroNodeV2(head, person2)
	InsertHeroNodeV2(head, person1)
	//InsertPersonNode(head, person1)
	//InsertPersonNode(head, person2)
	//InsertPersonNode(head, person3)
	//InsertPersonNode(head, person4)

	ListPersonNode(head)

	fmt.Println()
	DeleteNodeElement(head, 4)
	DeleteNodeElement(head, 2)
	ListPersonNode(head)

	fmt.Println()
	err := EditNodeElement(head, &PersonNode{
		no:   1,
		name: "Ezra Chai",
	})
	if err != nil {
		fmt.Println(err)
	}
	ListPersonNode(head)

}
