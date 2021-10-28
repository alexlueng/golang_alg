package amazon

// feature 4 we are integrating data from a recently acquired e-commerce website. Create a deep copy of
// a list of products from the acquired company's website so that it may be imported to Amazon.
// Solution
// this problem can be descripe as deep copy a complex linkedlist
// 1) traverse the linked list starting at the head
// 2) use a dictionary/map to keep track of visited nodes
// 3) make a copy of the current node and store the old node as the key and the new node as the value
// in visited
// 4) if the related pointer of the current node, i, points to the node j and a clone of j already
// exists in visited, we will use the cloned node as a reference.
// 5) otherwise, if the related pointer of the current node, i, points to the node j, which has not been
// created yet, we will create a new node that corresponds to j and add it to visited.
// 6) if the next pointer of the current node, i, points to the node j and a clone of j already exists in
// visited, we will use the cloned node as a reference.
// 7) if the next pointer of the current node, i, points to the node j, which has not been created yet,
// we will create a new node corresponding to j and add it to the visited dictionary.

type Node struct {
	prod    int
	next    *Node
	related *Node
}

func getClonedNode(node *Node, visited map[*Node]*Node) *Node {
	if node != nil {
		if _, ok := visited[node]; ok {
			return visited[node]
		} else {
			visited[node] = &Node{prod: node.prod}
			return visited[node]
		}
	}
	return nil
}

func copyProductRelations(head *Node, visited map[*Node]*Node) *Node {
	if head == nil {
		return head
	}
	oldNode := head

	newNode := &Node{prod: oldNode.prod}
	visited[oldNode] = newNode

	for oldNode != nil {
		newNode.related = getClonedNode(oldNode.related, visited)
		newNode.next = getClonedNode(oldNode.next, visited)

		oldNode = oldNode.next
		newNode = newNode.next
	}
	return visited[head]
}
