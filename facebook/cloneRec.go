package facebook

// Feature 2 We want all the user's friends on Facebook to be suggested on Instagram as well. Since
// Instagram is a different platform, all of its connections need to be copied to a separate database.

func cloneRec(root *Node, nodesCompleted map[*Node]*Node) *Node {
	if root == nil {
		return nil
	}
	newNode := &Node{data: root.data}
	nodesCompleted[root] = newNode

	for _, p := range root.friends {
		x := nodesCompleted[p]
		if x == nil {
			newNode.friends = append(newNode.friends, cloneRec(p, nodesCompleted))
		} else {
			newNode.friends = append(newNode.friends, x)
		}
	}
	return newNode
}

func clone(root *Node) *Node {
	nodesCompleted := make(map[*Node]*Node)
	return cloneRec(root, nodesCompleted)
}
