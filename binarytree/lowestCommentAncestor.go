package main

import "fmt"

// method 1: record the path of the two given target
// O(n) traversed the tree twice
func findPath(root *Treenode, path []int, target int) ([]int, bool) {
	// fmt.Println(path)
	if root == nil {
		return path, false
	}
	path = append(path, root.Val)
	if root.Val == target {
		return path, true
	}

	if lpath, ok := findPath(root.Left, path, target); ok {
		return lpath, ok
	}
	if rpath, ok := findPath(root.Right, path, target); ok {
		return rpath, ok
	}

	path = path[:len(path)-1]
	return path, false
}

func findLCA(root *Treenode, t1, t2 int) int {
	path := []int{}
	n1path, ok1 := findPath(root, path, t1)
	n2path, ok2 := findPath(root, path, t2)
	if !ok1 || !ok2 {
		return -1
	}
	fmt.Println(n1path)
	fmt.Println(n2path)
	var i int
	for i < len(n1path) && i < len(n2path) {
		if n1path[i] != n2path[i] {
			break
		}
		i++
	}
	return n1path[i-1]

}

// method 2: traversal one time
// assume both node are in the tree
func findLCAByOneTraversal(root *Treenode, v1, v2 int) *Treenode {
	if root == nil {
		return nil
	}
	if root.Val == v1 || root.Val == v2 {
		return root
	}
	left_lca := findLCAByOneTraversal(root.Left, v1, v2)
	right_lca := findLCAByOneTraversal(root.Right, v1, v2)
	if left_lca != nil && right_lca != nil {
		return root
	}
	if left_lca != nil && right_lca == nil {
		return left_lca
	}
	return right_lca
}
