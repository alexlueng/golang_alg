package amazon

import "sort"

// merge product recommendations using the data from an acquired company.
// Solution
// this can be mapped to a graph problem.

func accountsMerge(accounts [][]string) [][]string {
	emailToName := make(map[string]string)
	graph := make(map[string]map[string]bool)

	for _, acc := range accounts {
		name := acc[0]
		for i := 1; i < len(acc); i++ {
			email := acc[i]
			if _, ok := graph[acc[1]]; !ok {
				graph[acc[1]] = make(map[string]bool)
			}
			graph[acc[1]][email] = true
			if _, ok := graph[email]; !ok {
				graph[email] = make(map[string]bool)
			}
			graph[email][acc[1]] = true
			emailToName[email] = name
		}
	}

	seen := make(map[string]bool)
	var ans [][]string
	for g, _ := range graph {
		email := g
		if _, ok := seen[email]; !ok {
			seen[email] = true
			var stack Stack
			stack = stack.Push(email)
			var component []string
			for !stack.Empty() {
				node := stack.Top()
				stack, _ = stack.Pop()
				component = append(component, node)
				for nei, _ := range graph[node] {
					seen[nei] = true
					stack = stack.Push(nei)
				}
			}
			sort.Strings(component)
			component = append([]string{emailToName[email]}, component...)
			ans = append(ans, component)
		}
	}
	return ans
}
