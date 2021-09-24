package facebook

import (
	"fmt"
	"math/rand"
)

type Pair struct {
	first, second int
}

// if there is an edge from x to y
// that means there must be an edge from y to x
// and there is no edge form a node to itself
// hence there can maximim of (nodes * nodes - nodes) / 2 edges in this graph
func CreateTestGraphDirected(nodesCount, edgesCount int) []*Node {
	var vertices []*Node
	for i := 0; i < nodesCount; i++ {
		vertices = append(vertices, &Node{data: i})
	}
	var allEdges []Pair
	for i := 0; i < len(vertices); i++ {
		for j := i + 1; j < len(vertices); j++ {
			allEdges = append(allEdges, Pair{first: i, second: j})
		}
	}

	rand.Shuffle(len(allEdges), func(i, j int) { allEdges[i], allEdges[j] = allEdges[j], allEdges[i] })

	for i := 0; i < edgesCount && i < len(allEdges); i++ {
		edge := allEdges[i]
		vertices[edge.first].friends = append(vertices[edge.first].friends, vertices[edge.second])
		vertices[edge.second].friends = append(vertices[edge.second].friends, vertices[edge.second])
	}

	return vertices
}

func PrintGraph1(vertices []*Node) {
	for _, n := range vertices {
		fmt.Printf("%v: {", n.data)
		for _, t := range n.friends {
			fmt.Printf("%v ", t.data)
		}
		fmt.Println("}")
	}
}

func PrintGraph2(root *Node, visitedNodes map[*Node]bool) {
	if _, ok := visitedNodes[root]; root == nil || ok {
		return
	}
	visitedNodes[root] = true
	fmt.Printf("%v {", root.data)
	for _, n := range root.friends {
		fmt.Printf("%v ", n.data)
	}
	fmt.Println("}")
	for _, n := range root.friends {
		PrintGraph2(n, visitedNodes)
	}
}

func PrintGraph(root *Node) {
	visitedNodes := make(map[*Node]bool)
	PrintGraph2(root, visitedNodes)
}

func AreGraphsEqualRec(root1 *Node, root2 *Node, visited map[*Node]bool) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 != nil || root2 != nil {
		return false
	}
	if root1.data != root2.data {
		return false
	}
	if len(root1.friends) != len(root2.friends) {
		return false
	}
	for _, nbr1 := range root1.friends {
		found := false
		for _, nbr2 := range root2.friends {
			if nbr1.data == nbr2.data {
				if _, ok := visited[nbr1]; ok {
					visited[nbr1] = true
					AreGraphsEqualRec(nbr1, nbr2, visited)
				}
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}
