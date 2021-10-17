package binarytree

import "fmt"

/*
Depth First Traversals:
(a) Inorder (Left, Root, Right) : 4 2 5 1 3
(b) Preorder (Root, Left, Right) : 1 2 4 5 3
(c) Postorder (Left, Right, Root) : 4 5 2 3 1
*/

func inorder(root *Treenode) {
	if root == nil {
		return
	}
	inorder(root.Left)
	fmt.Printf("%d ", root.Val)
	inorder(root.Right)
}

func preorder(root *Treenode) {
	if root == nil {
		return
	}
	fmt.Printf("%d ", root.Val)
	preorder(root.Left)
	preorder(root.Right)
}

func postorder(root *Treenode) {
	if root == nil {
		return
	}
	postorder(root.Left)
	postorder(root.Right)
	fmt.Printf("%d ", root.Val)
}
