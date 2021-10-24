package uber

import "math"

// feature 2 calculate the max depth of water along the way between two points in a rain-affected city
// Solution
// water = minimum(leftMax, rightMax) - valueX
// 1) traverse through the list once
// 2) calculate the leftmax for each point and save it in ta list
// 3) traverse the list again to the same with rightmax
// 4) traverse the list once more, and use this data to find the amount of water above each point using
// the formula mentioned above.

func pathCost(elevationMap []int) int {
	water := 0
	n := len(elevationMap)

	leftMax := make([]int, n)
	rightMax := make([]int, n)

	leftMax[0] = elevationMap[0]
	rightMax[0] = elevationMap[0]

	for i := 1; i < n; i++ {
		leftMax[i] = int(math.Max(float64(leftMax[i-1]), float64(elevationMap[i])))
	}
	for i := n - 2; i >= 0; i-- {
		leftMax[i] = int(math.Max(float64(leftMax[i+1]), float64(elevationMap[i])))
	}
	for i := 0; i < n; i++ {
		water += int(math.Min(float64(leftMax[i]), float64(rightMax[i]))) - elevationMap[i]
	}
	return water
}
