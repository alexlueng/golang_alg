package amazon

// feature 9 allow a user to set a minimum and maximum price range filter to the list of products.
// Solution
// using a variation of the preorder traversal on the binary tree

func preOrder(node *BNode, low, high int, output []int) []int {
	if node != nil {
		if node.val <= high && low <= node.val {
			output = append(output, node.val)
		}
		if low <= node.val {
			output = preOrder(node.leftChild, low, high, output)
		}
		if node.val <= high {
			output = preOrder(node.rightChild, low, high, output)
		}
	}
	return output
}

func productsInRange(root *BNode, low, high int) []int {
	var output []int
	return preOrder(root, low, high, output)
}
