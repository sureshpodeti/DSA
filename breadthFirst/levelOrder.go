package breadthFirst

import (
	"container/list"
	"fmt"
)

/*
b := &BinaryTree{}

	ar := []int{12, 7, 9, 15, 13}

	for _, v := range ar {
		b.Add(v)
	}
	b.levelOrder()
*/

type Node struct {
	Data        int
	Left, Right *Node
}

type BinaryTree struct {
	Root *Node
}

func (b *BinaryTree) Add(value int) {
	b.Root = b.insertRecursive(b.Root, value)
}

func (b *BinaryTree) insertRecursive(root *Node, value int) *Node {
	if root == nil {
		return &Node{Data: value}
	}
	if value < root.Data {
		root.Left = b.insertRecursive(root.Left, value)
	} else {
		root.Right = b.insertRecursive(root.Right, value)
	}
	return root
}

func (b *BinaryTree) levelOrder() {
	queue := list.New()
	var q []any

	queue.PushBack(b.Root)
	for queue.Len() != 0 {
		size := queue.Len()
		for i := 0; i < size; i++ {
			element := queue.Front()
			node := element.Value.(*Node)
			queue.Remove(element)
			q = append(q, node.Data)

			if node.Left != nil {
				queue.PushBack(node.Left)
			}

			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}

	}

	fmt.Println(q)

}
