package amazon

// feature 7  help the logistic division of the company to optimally utilize contractor delivery trucks

func checkDelivery(packages []int, k int) bool {
	currSum := 0
	dict := make(map[int]int)
	dict[0] = -1
	for i := 0; i < len(packages); i++ {
		currSum += packages[i]
		if k != 0 {
			currSum = currSum % k
		}
		if _, ok := dict[currSum]; ok {
			if i-dict[currSum] > 1 {
				return true
			}
		} else {
			dict[currSum] = i
		}
	}
	return false
}
