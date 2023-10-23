package in_place_linked_list_reversal

import "fmt"

/*
l := &in_place_linked_list_reversal.LinkedList{}

	//for i := 0; i < 11; i++ {
	//	l.add(i)
	//}

	ar := []int{2, 4, 6, 8, 10}

	for _, val := range ar {
		l.Add(val)
	}
	l.Reverse()
	l.Display()
*/
type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func (l *LinkedList) Add(value int) {
	l.Head = l.insertRecursive(l.Head, value)

}
func (l *LinkedList) insertRecursive(head *Node, value int) *Node {
	if head == nil {
		return &Node{Data: value}
	}
	head.Next = l.insertRecursive(head.Next, value)
	return head
}

func (l *LinkedList) Display() {
	curr := l.Head

	for curr != nil {
		fmt.Println(curr.Data)
		curr = curr.Next
	}
}

func (l *LinkedList) Reverse() {
	var prev *Node
	curr := l.Head

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	l.Head = prev
}
