package binarytree

import "fmt"

func findMaxSumPath(root *Treenode) {
	sum := getRootToLeafSum(root)
	fmt.Println("the maximum sum is: ", sum)
	fmt.Println("the maximum path is: ")
	printPath(root, sum)
}

// Function to calculate the maximum root-to-leaf sum in a binary tree
func getRootToLeafSum(root *Treenode) int {
	if root == nil {
		return 0
	}
	left := getRootToLeafSum(root.Left)
	right := getRootToLeafSum(root.Right)
	return max(left, right) + root.Val
}

// Function to print the root-to-leaf path with a given sum in a binary tree
func printPath(root *Treenode, sum int) bool {
	if sum == 0 {
		return true
	}
	if root == nil {
		return false
	}
	left := printPath(root.Left, sum-root.Val)
	right := printPath(root.Right, sum-root.Val)

	if left || right {
		fmt.Print(root.Val)
	}
	return left || right
}
