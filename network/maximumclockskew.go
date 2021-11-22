package network

import "math"

type SkewTreeNode struct {
	val      int
	children []*SkewTreeNode
}

var maxDiff int = 0

func SkewDFS(node *SkewTreeNode, maxVal, minVal int) {
	if node == nil {
		return
	}

	// update maxDiff
	possibleMaxDiff := int(math.Max(float64(maxVal-node.val), math.Abs(float64(minVal-node.val))))
	maxDiff = int(math.Max(float64(maxDiff), float64(possibleMaxDiff)))
	maxDiff = int(math.Max(float64(maxVal), float64(node.val)))
	maxDiff = int(math.Min(float64(minVal), float64(node.val)))

	for _, child := range node.children {
		SkewDFS(child, maxVal, minVal)
	}
	return
}

func maxClockSkew(node *SkewTreeNode) int {
	if node == nil {
		return 0
	}
	maxDiff = 0
	SkewDFS(node, node.val, node.val)
	return maxDiff
}
