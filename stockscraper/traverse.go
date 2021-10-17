package stockscraper

import (
	"container/list"
)

// feature 1 develop a way to parse the DOM tree structure of different stock websites.
// Solution
// using breadth first search

type TreeNode struct {
	val      string
	children []*TreeNode
}

func traversal(root *TreeNode) [][]string {
	var result [][]string
	if root == nil {
		return nil
	}

	queue := list.New()
	queue.PushBack(root)
	for queue.Len() != 0 {
		levelSize := queue.Len()
		var currentLevel []string
		for i := 0; i < levelSize; i++ {
			temp := queue.Front()
			queue.Remove(temp)
			currentNode := temp.Value.(*TreeNode)
			currentLevel = append(currentLevel, currentNode.val)
			for _, child := range currentNode.children {
				queue.PushBack(child)
			}
		}
		result = append(result, currentLevel)
	}
	return result
}
