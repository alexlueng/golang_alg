package main

import "fmt"

// td p6
func getAncestorRecursive(root *Treenode, target int) bool {
	if root == nil {
		return false
	}
	if root.Val == target {
		return true
	}

	left := getAncestorRecursive(root.Left, target)

	right := false
	if !left {
		right = getAncestorRecursive(root.Right, target)
	}
	if left || right {
		fmt.Println("the ancestor is: ", root.Val)
	}
	return left || right
}
