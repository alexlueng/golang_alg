package network

import "math"

// feature 7 distribute files over a network to high-performance cluster nodes to minimize communication overhead.
// Solution
// 1) Compute the last occurrence of each letter in the string.
// 2) Initialize the start and end pointers to zero to denote the start of the string. Also, initialize a
// count variable to zero
// 3) keep updating the end pointer to the maximum last occurrence value for each character until we reach
// end during the traversal
// 4) Increment the value of count by 1 when the above condition is fulfilled
// 5) assign the start pointer to end+1 to denote the start of the next split

func partitionLabels(files string) int {
	last := make([]int, 26)
	for i := 0; i < len(files); i++ {
		last[files[i]-'a'] = i
	}

	end, count := 0, 0
	for i := 0; i < len(files); i++ {
		end = int(math.Max(float64(end), float64(last[files[i]-'a'])))
		if i == end {
			count++
		}
	}
	return count
}
