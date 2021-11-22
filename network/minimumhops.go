package network

import "math"

// feature 3 Each router in a linear virtual network topology can send a packet a certain maximum
// hops forward. We need to determine the minimum number of transmissions required to get a packet
// from the first router to the last.
// Solution
// 1) init the maximum position that one could reach and the maximum position reachable during the
// first hub as maxReach = nums[0] and currReach = nums[0] respectively
// 2) if the size of our list is greater than 1, then initialize the hops variable to 1
// 3) start iterating over the array
// 4) in each iteration, update the maxReach so it contains the maximum position one could reach from
// the current index i as max of maxReach and i + value[i]
// 5) in each iteration, if currReach < i, it means that we are ready for another hop,.increment hop
// by 1 and assign currReach to the current maximum position which is maxReach
// 6) return the number of hops when the loop exits

func minimumHops(values []int) int {
	if len(values) < 2 {
		return 0
	}
	maxReach := values[0]
	currReach := values[0]

	hops := 1
	for i := 0; i < len(values); i++ {
		if currReach < i {
			hops++
			currReach = maxReach
		}
		maxReach = int(math.Max(float64(maxReach), float64(values[i]+i)))
	}
	return hops
}
