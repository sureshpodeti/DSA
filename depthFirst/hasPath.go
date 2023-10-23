package depthFirst

/*
	b := &BinaryTree{}

	ar := []int{10, 7, 5, 8, 13, 12, 15}

	S := 35

	for _, v := range ar {
		b.Add(v)
	}

	hasPath := b.hasPath(S)

	fmt.Printf("Found path: %v\n", hasPath)
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

func (b *BinaryTree) hasPath(S int) bool {
	return b.hasPathUtil(b.Root, S)
}
func (b *BinaryTree) hasPathUtil(root *Node, S int) bool {
	if root == nil {
		return false
	}

	if root.Data == S && root.Left == nil && root.Right == nil {
		return true
	}

	return b.hasPathUtil(root.Left, S-root.Data) || b.hasPathUtil(root.Right, S-root.Data)

}
