package zoom

import "strings"

// feature 2 participant data needs to be serialized from the server to the network and de-serialized
// at the client.
// Solution
// this can be descripted as the serialize a binary tree

func serialize(root *Node) string {
	var res []string
	str := ""
	res = preorder(root, res)
	for i := 0; i < len(res); i++ {
		str += res[i]
		if i+1 < len(res) {
			str += ","
		}
	}
	return str
}

func preorder(root *Node, res []string) []string {
	if root != nil {
		res = append(res, root.val)
		res = preorder(root.leftChild, res)
		res = preorder(root.rightChild, res)
	}
	return res
}

func deserialized(data string) *Node {
	var root *Node
	delim := ","

	lst := strings.Split(data, delim)

	var stack Stack
	for _, name := range lst {
		if root == nil {
			root = &Node{val: name}
			stack = stack.Push(root)
		} else {
			root.Insert(name)
		}
	}
	return root
}
