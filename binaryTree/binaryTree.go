package binaryTree

import "fmt"

/*
	b := &binaryTree.BinaryTree{}

	ar := []int{2, 1, 3}

	for _, val := range ar {
		b.Add(val)
	}

//b.PreOrder()
//b.InOrder()
//b.PostOrder()
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

func (b *BinaryTree) PreOrder() {
	b.PreOrderRecursive(b.Root)
}
func (b *BinaryTree) PreOrderRecursive(root *Node) {
	if root == nil {
		return
	}
	fmt.Println(root.Data)
	b.PreOrderRecursive(root.Left)
	b.PreOrderRecursive(root.Right)
}

func (b *BinaryTree) InOrder() {
	b.InOrderRecursive(b.Root)
}

func (b *BinaryTree) InOrderRecursive(root *Node) {
	if root == nil {
		return
	}
	b.InOrderRecursive(root.Left)
	fmt.Println(root.Data)
	b.InOrderRecursive(root.Right)
}

func (b *BinaryTree) PostOrder() {
	b.PostOrderRecursive(b.Root)
}
func (b *BinaryTree) PostOrderRecursive(root *Node) {
	if root == nil {
		return
	}
	b.PostOrderRecursive(root.Left)
	b.PostOrderRecursive(root.Right)
	fmt.Println(root.Data)
}
