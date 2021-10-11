package searchengine

// feature 5 calculate a search ranking factor based on the ranking score of pages that refer to it

// Solution
// The optimal approach for solving this problem is that for every index i, we will evaluate the product
// of the numbers to the left and all the numbers to the right of i.

func searchRanking(pageScores []int) []int {
	length := len(pageScores)
	ranking := make([]int, length)
	ranking[0] = 1
	for i := 1; i < length; i++ {
		ranking[i] = pageScores[i-1] * ranking[i-1]
	}

	right := 1
	for i := length - 1; i >= 0; i-- {
		ranking[i] = ranking[i] * right
		right *= pageScores[i]
	}
	return ranking
}
