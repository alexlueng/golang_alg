package calender

// feature 3 before scheduling a meeting, check if the other person will be availbale during the
// specified time slot
// we want to program a function that will let user A know if it is possible to schedule a meeting
// with user B or not.
// Solution
// 1) use brute force traversal
// 2) use binary search tree. we can insert all the meetings in a BST first, and then check if the new
// meeting can also be inserted without any clash.

type Node struct {
	start      int
	end        int
	leftChild  *Node
	rightChild *Node
}

type BST struct {
	root *Node
}

func (b *BST) Insert(start, end int) bool {
	if b.root == nil {
		b.root = &Node{start: start, end: end}
		return true
	}
	newNode := &Node{start: start, end: end}
	return b.AddNode(b.root, newNode)
}

func (b *BST) AddNode(currentNode, newNode *Node) bool {
	if newNode.start >= currentNode.end {
		if currentNode.rightChild == nil {
			currentNode.rightChild = newNode
			return true
		}
		return b.AddNode(currentNode.rightChild, newNode)
	} else if newNode.end <= currentNode.start {
		if currentNode.leftChild == nil {
			currentNode.leftChild = newNode
			return true
		}
		return b.AddNode(currentNode.leftChild, newNode)
	}
	return false
}

func checkMeeting(meetingTimes [][]int, newMeeting []int) bool {
	tree := &BST{}
	for _, meeting := range meetingTimes {
		tree.Insert(meeting[0], meeting[1])
	}
	return tree.Insert(newMeeting[0], newMeeting[1])
}
