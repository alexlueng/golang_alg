package zoom

// feature 3 implement an algorithm to find the correct answer of minigame that prompts the user to
// find the minimum number of steps taken to reach from one point to another
// Solution
// THis problem can be mapped as a graph problem in which we need to find the shortest path between two
// vertices.

func minStemps(k []int) int {
	n := len(k)
	if n <= 1 {
		return 0
	}

	graph := make(map[int][]int)
	for i := 0; i < n; i++ {
		if _, ok := graph[k[i]]; ok {
			graph[k[i]] = []int{i}
		} else {
			graph[k[i]] = append(graph[k[i]], i)
		}
	}

	current := []int{0}

	visited := make(map[int]bool)
	step := 0

	for len(current) != 0 {
		var nextNode []int
		for _, node := range current {
			if node == n-1 {
				return step
			}
			for _, child := range graph[k[node]] {
				if _, ok := visited[child]; !ok {
					visited[child] = true
					nextNode = append(nextNode, child)
				}
			}
			graph[k[node]] = []int{}

			if _, ok := visited[node+1]; !ok && node+1 < n {
				visited[node+1] = true
				nextNode = append(nextNode, node+1)
			}
			if _, ok := visited[node-1]; !ok && node-1 >= 0 {
				visited[node-1] = true
				nextNode = append(nextNode, node-1)
			}
		}
		current = nextNode
		step++
	}
	return -1
}
