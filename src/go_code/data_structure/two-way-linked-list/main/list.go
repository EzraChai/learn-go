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
			temp.LastNode.BeforeNode = newNode
			newNode.LastNode = temp.LastNode
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
	fmt.Println(temp)
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
		fmt.Println(*temp)
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
				fmt.Println(*temp)
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
	DeleteNode(headNode, 1)
	fmt.Println("***********")
	ListNode(headNode)

}
