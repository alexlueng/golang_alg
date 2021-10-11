package searchengine

// Feature 3 check if white-spaces can be added to a query to create valid words in case the original
// query does not get any hits on the web
// this module will be triggered if the original query gives very few results or no results at all

// we use dynamic programming technique.
// Solution
// 1) create a boolean list called dp to track all the substrings in the query that satisfy the given condition
// 2) use two nested loops with i and j as indices. The index i will refer to the length of the substring pi and j will be the index that splits the current substring pj into smaller substrings.
// 3) initially, the first element of the dp array will be set to true because an empty string is always
// present in the dictionary, and the rest of the elements of the dp are set to false.
// 4) in the nested loop, the substrings of all possible lengths are considered starting from the beginning
// by using the index i. Each substring is then split into p1 and p2 in every possible way using the index j.
// 5) next, the entry at dp[i] is checked to see if it contains True, which means that the substring p1
// satisfies the given condition. If this condition is satisfied, we further check if p2 is also present in the dictionary.
// 6) if both strings are present, we set dp[i+1] to True. Otherwise, it remains False. Here l represents the length of the substring.

func breakQuery(query string, dict []string) bool {
	dictSet := make(map[string]bool)
	for _, val := range dict {
		dictSet[val] = true
	}
	dp := make([]bool, len(query)+1)
	dp[0] = true
	for i := 1; i <= len(query); i++ {
		for j := 0; j < i; j++ {
			if _, ok := dictSet[query[j:i]]; ok && dp[j] {
				dp[i] = true
			}
		}
	}
	return dp[len(query)]
}
