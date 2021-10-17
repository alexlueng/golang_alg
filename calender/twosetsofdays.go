package calender

import "math"

// feature 6 find two sets of consecutive days to schedule a series of board meetings using the number
// of mutually free hours of the board members.

func twoSetsOfDays(hoursPerDay []int, k int) int {
	hmap := make(map[int]int)
	sum := 0
	lsize := math.MaxInt64
	result := math.MaxInt64
	hmap[0] = -1
	for i := 0; i < len(hoursPerDay); i++ {
		sum += hoursPerDay[i]
		hmap[sum] = i
	}

	sum = 0
	for i := 0; i < len(hoursPerDay); i++ {
		sum += hoursPerDay[i]
		if val, ok := hmap[sum-k]; ok {
			// stores minimum length of sub-array ending with index <= i with sum k.
			// this ensures non-overlapping property
			lsize = int(math.Min(float64(lsize), float64(i-val)))
		}
		// searches for any sub-array startwith index i+1 with sum k
		if val, ok := hmap[sum+k]; ok && lsize < math.MaxInt64 {
			result = int(math.Min(float64(result), float64(val-i+lsize)))
		}
	}

	if result == math.MaxInt64 {
		return -1
	}
	return result
}
