package main

import "fmt"

func main() {

	root := Treenode{Val: 1}
	root.Left = &Treenode{Val: 2}
	root.Right = &Treenode{Val: 3}
	root.Left.Left = &Treenode{Val: 8}
	root.Left.Right = &Treenode{Val: 4}
	root.Right.Left = &Treenode{Val: 5}
	root.Right.Right = &Treenode{Val: 6}
	root.Left.Right.Left = &Treenode{Val: 10}
	root.Right.Left.Left = &Treenode{Val: 7}
	root.Right.Left.Right = &Treenode{Val: 9}
	root.Right.Right.Right = &Treenode{Val: 12}

	fmt.Println("inorder traversal: ")
	inorder(&root)
	fmt.Println("\n---------------")

	fmt.Println("preorder traversal: ")
	preorder(&root)
	fmt.Println("\n---------------")

	fmt.Println("Construct trees: ")
	in := []int{8, 2, 10, 4, 1, 7, 5, 9, 3, 6, 12}
	pre := []int{1, 2, 8, 4, 10, 3, 5, 7, 9, 6, 12}

	preIndex := 0
	tree1 := buildTreeFromPreAndIn(pre, in, &preIndex, 0, len(in)-1)
	fmt.Println("preorder traversal the new tree: ")
	preorder(tree1)
	fmt.Println()

	tree2 := buildTreeWrap(pre, in)
	fmt.Println("preorder traversal the new tree using map=: ")
	preorder(tree2)
	// findMaxSumPath(&root)
	// fmt.Println("\n---------------")
	// getAncestorRecursive(&root, 12)

	// lca := findLCA(&root, 9, 7)
	// fmt.Println("***the lca is: ", lca)
	// lca_node := findLCAByOneTraversal(&root, 10, 7)
	// fmt.Println("&&&the lca is: ", lca_node.Val)
}
