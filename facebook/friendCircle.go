package facebook

// feature 1 find all the people on Facebook that are in a user's friend circle.
// Your individual friend circle is defined as a group of users who are directly or indirectly
// friends with you.
// The input is an n*n square matrix
// Solution
// Think of the symmetric input matrix as an undirected graph. All the friends who should be in one friend
// circle are also in one connected component in the graph.
// 1) Initialize a list/array, called visited, to keep track of the visited vertices of size n with 0
// as the initial value at each index.
// 2) For every vertex v, traverse the graph using DFS if visited[v] is 0. Otherwise, move to the next v.
// 3) Set visited[v] to 1 for every v that the DFS traversal encounters.
// 4) When the DFS traversal terminates, increment the friend circles counter by 1. This means that one whole connected component has been traversed.

// input example:
//   A B C D
// A % % 0 0
// B % % % 0
// C 0 % % 0
// D 0 0 0 %
func DFS(friends [][]bool, n int, visited []bool, v int) {
	for i := 0; i < n; n++ {
		if friends[v][i] && !visited[i] && i != v {
			visited[i] = true
			DFS(friends, n, visited, i)
		}
	}
}

func friendCircles(friends [][]bool, n int) int {
	if n == 0 {
		return 0
	}
	numCircles := 0
	visited := make([]bool, n)
	for i := 0; i < n; i++ {
		if !visited[i] {
			visited[i] = true
			DFS(friends, n, visited, i)
			numCircles++
		}
	}
	return numCircles
}
