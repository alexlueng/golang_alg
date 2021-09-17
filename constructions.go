package main

// construct from preorder and inorder
// inorder traversal:
// 8 2 10 4 1 7 5 9 3 6 12
// preorder traversal:
// 1 2 8 4 10 3 5 7 9 6 12
// from preorder we can see 1 is the root
// then we can search 1 from inorder. and the we can tell the left subtree and right subtree
// we can do this recursively

// Algorithm: buildTree()
// 1) Pick an element from Preorder. Increment a Preorder Index Variable (preIndex in below code) to pick the next element in the next recursive call.
// 2) Create a new tree node tNode with the data as the picked element.
// 3) Find the picked element’s index in Inorder. Let the index be inIndex.
// 4) Call buildTree for elements before inIndex and make the built tree as a left subtree of tNode.
// 5) Call buildTree for elements after inIndex and make the built tree as a right subtree of tNode.
// 6) return tNode.

// Time complexity: O(n2)

func search(arr []int, start, end, target int) int {
	for i := start; i <= end; i++ {
		if arr[i] == target {
			return i
		}
	}
	return -1
}

func buildTreeFromPreAndIn(pre, in []int, preIndex *int, start, end int) *Treenode {
	// fmt.Printf("start %d, end %d\n", start, end)
	if start > end {
		return nil
	}
	// preIndex := 0
	// fmt.Println("preIndex: ", *preIndex)
	if *preIndex > len(pre) {
		return nil
	}
	root := &Treenode{Val: pre[*preIndex]}
	*preIndex += 1

	if start == end {
		return root
	}
	inIndex := search(in, start, end, root.Val)
	root.Left = buildTreeFromPreAndIn(pre, in, preIndex, start, inIndex-1)
	root.Right = buildTreeFromPreAndIn(pre, in, preIndex, inIndex+1, end)
	return root
}

// effective approach：store indexes of inorder traversal in a hashtable

func buildTreeUsingMap(pre, in []int, preIndex *int, start, end int, mp map[int]int) *Treenode {
	if start > end {
		return nil
	}
	if *preIndex > len(pre) {
		return nil
	}
	root := &Treenode{Val: pre[*preIndex]}
	*preIndex += 1
	if start == end {
		return root
	}
	inIndex := mp[root.Val]
	root.Left = buildTreeFromPreAndIn(pre, in, preIndex, start, inIndex-1)
	root.Right = buildTreeFromPreAndIn(pre, in, preIndex, inIndex+1, end)
	return root
}

func buildTreeWrap(pre, in []int) *Treenode {
	mp := make(map[int]int)
	for i := 0; i < len(in); i++ {
		mp[in[i]] = i
	}
	preIndex := 0
	root := buildTreeFromPreAndIn(pre, in, &preIndex, 0, len(in)-1)
	return root
}
