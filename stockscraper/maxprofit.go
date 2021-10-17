package stockscraper

// feature 4 retrieve the stock increase and decrease percentages from the DOM tree and calculate the
// maximum profit we could obtain from it

func maxProfit(A []int) int {
	if len(A) < 1 {
		return 0
	}
	currMax := A[0]
	globalMax := 0

	for i := 1; i < len(A); i++ {
		if currMax < 0 {
			currMax = A[i]
		} else {
			currMax += A[i]
		}
		if globalMax < currMax {
			globalMax = currMax
		}
	}
	return globalMax
}
