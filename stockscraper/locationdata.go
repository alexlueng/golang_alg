package stockscraper

// feature 2 assign a confidence score to each HTML tag in the DOM to represent the likehood
// that the tag contains stock price data. Then filter the minimal subtree of the DOM that has stock
// price data.

// Solution
// The problem can be solved as Lowest Common Ancestor.
// 1) we traverse the tree starting from the root node
// 2) we store the parent of each node in the dictionary until the nodes a and b are not found
// 3) if the nodes a and b are found:
// 		a) traverse over the parents/ancestors of node a
// 		b) for each parent node, add it to the ancestors set
// similarly, we will traverse through the ancestors of node b. If the ancestor is present in the
// ancestors set fo a, this is the first ancestor common in between a and b(while traversing upwards),
// and this is the LCA node

type LCATreeNode struct {
	val      int
	children []*LCATreeNode
}

func lowerCommonAncestor(root, a, b *LCATreeNode) int {
	var stack Stack
	parent := make(map[*LCATreeNode]*LCATreeNode)

	parent[root] = nil
	stack = stack.Push(root)
	_, temp := parent[b]
	_, ok := parent[a]
	for !ok || !temp {
		node := stack.Top()
		stack, _ = stack.Pop()

		for _, child := range node.children {
			parent[child] = node
			stack = stack.Push(child)
		}
		_, temp = parent[b]
		_, ok = parent[a]
	}
	ancestors := make(map[*LCATreeNode]bool)
	for a != nil {
		ancestors[a] = true
		a = parent[a]
	}

	_, ok = ancestors[b]
	for !ok {
		b = parent[b]
		_, ok = ancestors[b]
	}
	return b.val
}
