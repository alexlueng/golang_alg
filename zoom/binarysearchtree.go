package zoom

type BinarySearchTree struct {
	root *Node
}

func (b *BinarySearchTree) Insert(val string) {
	if b.root == nil {
		b.root = &Node{val: val}
	} else {
		b.root.Insert(val)
	}
}
