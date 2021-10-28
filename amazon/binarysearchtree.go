package amazon

type BNode struct {
	val        int
	leftChild  *BNode
	rightChild *BNode
}

func (b *BNode) Insert(val int) {
	current := b
	parent := current
	for current != nil {
		parent = current
		if val < current.val {
			current = current.leftChild
		} else {
			current = current.rightChild
		}
	}
	if val < parent.val {
		parent.leftChild = &BNode{val: val}
	} else {
		parent.rightChild = &BNode{val: val}
	}
}

type BinarySearchTree struct {
	root *BNode
}

func (b *BinarySearchTree) Insert(val int) {
	if b.root == nil {
		b.root = &BNode{val: val}
	} else {
		b.root.Insert(val)
	}
}
