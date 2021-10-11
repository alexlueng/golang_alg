package searchengine

import "strings"

// feature 4 As an extension of the previous feature, find all the possible queries that can be created
// by adding white spaces to the original query

// Solution
// using top-down dynamic programming.

// helper function populates an array including all substrings derived from s.
func helper(s string, dictSet map[string]bool, dictMap map[string][]string) []string {
	var res []string
	if _, ok := dictMap[s]; ok {
		return dictMap[s]
	}
	if len(s) == 0 {
		res = append(res, "")
		return res
	}

	for word, _ := range dictSet {
		if strings.Index(s, word) == 0 {
			subList := helper(s[len(word):], dictSet, dictMap)
			for _, sub := range subList {
				temp := ""
				if len(sub) != 0 {
					temp = " "
				}
				res = append(res, word+temp+sub)
			}
		}
	}
	dictMap[s] = res
	return res
}

func possibleBreakQuery(query string, dict []string) []string {
	dictSet := make(map[string]bool)
	for _, val := range dict {
		dictSet[val] = true
	}
	dictMap := make(map[string][]string)
	return helper(query, dictSet, dictMap)
}
