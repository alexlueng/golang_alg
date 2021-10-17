package stockscraper

import "fmt"

// feature 3 stock price data at the same level of the DOM tree is closely related. Develop a new
// way to efficiently parse the DOM tree structure by introducing a next pointer in each tree node.
// pointing to the next node in the same level.
// Solution
// 1) root's next poiter will always be null. We establish the next pointers of the children of the root.
// 2) find the next level the leftmost node to be the starting point.
// 3) maintain a curr pointer that will be used to traverse current level's nodes.
// 4) use a prev pointer to establish the next pointers using the curr pointer.

type NTreeNode struct {
	val      string
	next     *NTreeNode
	children []*NTreeNode
}

func traversingDomTree(root, prev, leftmost *NTreeNode) *NTreeNode {
	if root == nil {
		return root
	}
	leftmost = root
	curr := leftmost
	for leftmost != nil {
		// prev tracks the lastest node on the "next" level
		// curr tracks the lastest node on the current level
		prev = nil
		curr = leftmost
		leftmost = nil
		// iterate on the nodes of current level like a linked list
		for curr != nil {
			fmt.Println(curr.val)
			// if we found at least one node on the new level
			// setup its next pointer
			for _, child := range curr.children {
				if child != nil {
					if prev != nil {
						prev.next = child
					} else {
						leftmost = child
					}
				}
				prev = child
			}
			curr = curr.next
		}
	}
	return root
}
