package network

// feature 2 given that a specific time-to-live has been set in a message, find all the nodes on the
// network where the packet can travel from the current device
// Solution
// Combine DFS and BFS
// 1) init the adj list
// 2) perform DFS on the root node and add every node's parents and children to the adj list
// 3) mark the server node as the result and iterate ttl times over the adj list starting from the server node
// 4) during iteration, we will get all the parents and children nodes of the current node and mark them
// as the result
// 5) when the loop ends, you'll have the nodes that are ttl distance from the server as the result

type TreeNode struct {
	val      int
	children []*TreeNode
}

func dfs(parent *TreeNode, child *TreeNode, neighbors map[int][]int) {
	if child == nil {
		return
	}

	if parent != nil {
		if _, ok := neighbors[parent.val]; !ok {
			neighbors[parent.val] = []int{}
		}
		if _, ok := neighbors[child.val]; !ok {
			neighbors[child.val] = []int{}
		}
		neighbors[parent.val] = append(neighbors[parent.val], child.val)
		neighbors[child.val] = append(neighbors[child.val], child.val)
	}
	for i := 0; i < len(child.children); i++ {
		dfs(child, child.children[i], neighbors)
	}
}

func getDevices(root *TreeNode, server *TreeNode, ttl int) []int {
	neighbors := make(map[int][]int)
	dfs(nil, root, neighbors)

	var bfs []int
	bfs = append(bfs, server.val)
	lookup := make(map[int]bool)
	for _, i := range bfs {
		lookup[i] = true
	}
	for i := 0; i < ttl; i++ {
		var temp []int
		for _, node := range bfs {
			for _, nei := range neighbors[node] {
				if _, ok := lookup[nei]; !ok {
					temp = append(temp, nei)
				}
			}
		}
		bfs = temp
		for _, i := range bfs {
			lookup[i] = true
		}
	}
	return bfs
}
