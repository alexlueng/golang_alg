package uber

// feature 3 find a path from each driver's location to the user's location and calculate the cost of
// travel so we can choose the driver with the least cost.
// Solution
// using depth first search to find to a path between two nodes.
// 1) build the graph using the city map list Gmap
// 2) assign the cost to each edge while building the graph
// 3) once the graph is build, evaluate each driver's path in the drivers list by searching for a path
// between the driver node and the user node.
// 4) return the accumulated sum if the path exists. Otherwise, return -1

func backtrackEvaluate(city map[string]map[string]float64, currNode string, targetNode string, accSum float64, visited map[string]bool) float64 {
	visited[currNode] = true
	ret := -1.0

	neighbors := city[currNode]
	if i, ok := neighbors[targetNode]; ok {
		ret = accSum + i
	} else {
		for nextNode, val := range neighbors {
			if _, ok := visited[nextNode]; ok {
				continue
			}
			ret = backtrackEvaluate(city, nextNode, targetNode, accSum+val, visited)
			if ret != -1.0 {
				break
			}
		}
	}
	delete(visited, currNode)
	return ret
}

func getTotalCost(Gmap [][]string, pathCosts []float64, drivers []string, user string) []float64 {
	results := make([]float64, len(drivers))
	city := make(map[string]map[string]float64)

	// 1)
	for i := 0; i < len(Gmap); i++ {
		checkPoints := Gmap[i]
		sourceNode := checkPoints[0]
		destNode := checkPoints[1]
		pathCost := pathCosts[i]

		if _, ok := city[sourceNode]; !ok {
			city[sourceNode] = make(map[string]float64)
		}
		if _, ok := city[destNode]; !ok {
			city[destNode] = make(map[string]float64)
		}
		city[sourceNode][destNode] = pathCost
		city[destNode][sourceNode] = pathCost
	}

	// 2)
	for i := 0; i < len(drivers); i++ {
		driver := drivers[i]
		_, temp := city[user]
		if _, ok := city[driver]; !ok || !temp {
			results[i] = -1.0
		} else {
			visited := make(map[string]bool)
			results[i] = backtrackEvaluate(city, driver, user, 0, visited)
		}
	}
	return results
}
