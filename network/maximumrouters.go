package network

import "math"

// feature 4 disseminate an important packet to the maximum number of routers in a grid topology.
// subject to certain forwarding constraints.
// Solution
// Memoization DFS
// 1) init a res variable and assign it a value 0
// 2) recursively call DFS on each cell of the matrix, and update res if the result of the DFS call
// is greater the current res value
// 3) during a DFS call. keep visiting each of the four paths from the current cell if the value in
// the next cell is greater than the value in the current cell
// 4) if the result for a visited cell is not calculated, we compute it and cache(or memoize) it
// 5) otherwise, we get it from the cache directly
// 6) finally, return the res variable

func dfsRouters(grid [][]int, i, j, prevVal int, cache *[][]int) int {
	if i < 0 || j < 0 || i > len(grid)-1 || j > len(grid[0])-1 {
		return 0
	} else if prevVal > grid[i][j] {
		return 0
	} else if (*cache)[i][j] != 0 {
		return (*cache)[i][j]
	}
	// up
	pathUp := dfsRouters(grid, i-1, j, grid[i][j], cache)
	// down
	pathDown := dfsRouters(grid, i+1, j, grid[i][j], cache)
	// left
	pathLeft := dfsRouters(grid, i, j-1, grid[i][j], cache)
	// right
	pathRight := dfsRouters(grid, i, j+1, grid[i][j], cache)

	max1 := int(math.Max(float64(pathUp), float64(pathDown)))
	max2 := int(math.Max(float64(pathLeft), float64(pathRight)))

	(*cache)[i][j] = 1 + int(math.Max(float64(max1), float64(max2)))

	return (*cache)[i][j]
}

func maximumRouters(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	res := 0
	cache := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		cache[i] = make([]int, len(grid[0]))
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if cache[i][j] == 0 {
				prevVal := grid[i][j]
				dfsRouters(grid, i, j, prevVal, &cache)
				res = int(math.Max(float64(cache[i][j]), float64(res)))
			}
		}
	}
	return res
}
