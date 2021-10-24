package uber

import "math/rand"

// feature 5 implement the pool functionality in which drivers can pick up multiple customers. On picking
// up a pool customer, we will suggest a route to the driver on which they may find more pool customers.

type Uber struct {
	cumSums []int
}

func newUber(metrics []int) Uber {
	cumSum := 0
	uber := Uber{}
	for i := 0; i < len(metrics); i++ {
		cumSum += metrics[i]
		uber.cumSums = append(uber.cumSums, cumSum)
	}
	return uber
}

func (u Uber) PickRoute() int {
	random := rand.Float64()
	sum := float64(u.cumSums[len(u.cumSums)-1])
	target := random * sum

	low, high := 0, len(u.cumSums)
	for low < high {
		mid := low + (high-low)/2
		if target > float64(u.cumSums[mid]) {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return low
}
