package main

import "fmt"

type Node struct {
	No         int
	Name       string
	BeforeNode *Node //Point to the front node
	LastNode   *Node
}

func InsertHeroNode(head *Node, newNode *Node) {
	temp := head
	for {
		if temp.LastNode == nil {
			break
		}
		if temp.LastNode.No > newNode.No {
			newNode.LastNode = temp.LastNode
			temp.LastNode.BeforeNode = newNode
			break
		}
		temp = temp.LastNode
	}

	temp.LastNode = newNode
	newNode.BeforeNode = temp
}

func DeleteNode(head *Node, no int) {
	temp := head
	for {
		if temp.No == no {
			break
		}
		if temp.LastNode == nil {
			fmt.Println("Not found")
			return
		}

		temp = temp.LastNode
	}
	if temp.LastNode == nil {
		temp.BeforeNode.LastNode = nil
		return
	}
	temp.BeforeNode.LastNode = temp.LastNode
	temp.LastNode.BeforeNode = temp.BeforeNode
}

func ListNode(head *Node) {
	temp := head
	for {
		if temp.LastNode == nil {
			break
		}
		temp = temp.LastNode
		fmt.Printf("No: %v Name: %v, Current: %p, Before: %p, Last: %p\n", temp.No, temp.Name, temp, temp.BeforeNode, temp.LastNode)

	}
}

func ListReverseNode(head *Node) {
	temp := head
	for {
		if temp.LastNode == nil {
			for {
				if temp.BeforeNode == nil {
					break
				}
				fmt.Printf("No: %v Name: %v, Current: %p, Before: %p, Last: %p\n", temp.No, temp.Name, temp, temp.BeforeNode, temp.LastNode)
				temp = temp.BeforeNode
			}
			break
		}
		temp = temp.LastNode
	}
}

func main() {
	headNode := &Node{
		No:   0,
		Name: "",
	}
	InsertHeroNode(headNode, &Node{
		No:   1,
		Name: "Chloe Gan",
	})
	InsertHeroNode(headNode, &Node{
		No:   5,
		Name: "Ezra Chai",
	})
	InsertHeroNode(headNode, &Node{
		No:   3,
		Name: "Zoewin Tan",
	})
	ListReverseNode(headNode)
	fmt.Println("***********")
	ListNode(headNode)
	DeleteNode(headNode, 5)
	fmt.Println("***********")
	ListNode(headNode)

}
