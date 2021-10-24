package uber

import "math"

// feature 6 introduce the longest route of the city that will maximize the possibility of finding
// customers
// Solution
// this problem can be descripe as finding the longest path in a binary tree.

type LTreeNode struct {
	val   string
	left  *LTreeNode
	right *LTreeNode
}

func height(node *LTreeNode) int {
	if node == nil {
		return 0
	}
	lh := height(node.left)
	rh := height(node.right)

	return int(math.Max(float64(lh), float64(rh))) + 1
}

func longestRoute(root *LTreeNode) int {
	if root == nil {
		return 0
	}
	lHeight := height(root.left)
	rHeight := height(root.right)

	lpath := longestRoute(root.left)
	rpath := longestRoute(root.right)

	temp := math.Max(float64(lpath), float64(rpath))
	res := int(math.Max(float64(lHeight+rHeight+1), temp))

	return res
}
