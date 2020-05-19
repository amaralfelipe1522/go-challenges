package main

import (
	"fmt"
)

//Node is bla bla
type Node struct {
	num  int
	next *Node
}

// List is bla bla
type List struct {
	length int
	start  *Node
}

func main() {
	items := &List{}
	size := 10

	for i := 0; i <= size; i++ {
		node := Node{num: i}
		fmt.Println("Node (i):", node)
		items.insert(&node)
	}
	//fmt.Println("Items:", items)
	items.display()
}

func (l *List) insert(newNode *Node) {
	if l.length == 0 {
		l.start = newNode
	} else {
		currentNode := l.start
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = newNode
	}
	l.length++
}

func (l *List) display() {
	list := l.start
	for list.next != nil {
		fmt.Printf("%v ->\n", list.num)
		list = list.next
	}
	if list.next == nil {
		fmt.Printf("%v ->\n", list.num)
	}
}
