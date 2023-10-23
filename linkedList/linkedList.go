package linkedList

import "fmt"

/*
 linkedList := LinkedList{}

	for i := 0; i < 11; i++ {
		linkedList.insert(i)
	}

	linkedList.display()
*/

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func (l *LinkedList) insert(value int) {
	l.Head = l.insertRecursively(l.Head, value)

}

func (l *LinkedList) insertRecursively(head *Node, value int) *Node {
	if head == nil {
		return &Node{Data: value}
	}
	head.Next = l.insertRecursively(head.Next, value)
	return head
}

func (l *LinkedList) display() {
	curr := l.Head
	for curr != nil {
		fmt.Println(curr.Data)
		curr = curr.Next
	}
}
