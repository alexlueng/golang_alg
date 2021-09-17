package netflix

import "strconv"

// Feature 1 we want to enable users to see relevant search results despite minor typos

// "duel" "dule" "speed" "spede" "deul" "cars"
// split the list of titles into sets of words so that all words in a set are anagrams
// [{"dule", "duel", "deul"}, {"speed", "spede"}, {"cars"}]

// Solution
// all members of each set are characterized by the sam frequency of each alphabet
// 1 for each title, compute a 26-element vector abbccc #1#2#3#0#0#0...#0
// 2 use this vector as a key to insert into a hashmap
// 3 store the vector of the calculated character counts in the same hashmap as a key and assign the respective set of anagrams as its value
// 4 return the values of the hashmap.

func groupTitles(strs []string) [][]string {
	var output [][]string
	if len(strs) == 0 {
		return output
	}
	res := make(map[string][]string)
	for _, s := range strs {
		count := make([]int, 26)
		for _, c := range s {
			index := c - 'a'
			count[index]++
		}
		key := ""
		for i := 0; i < 26; i++ {
			key += "#"
			key += strconv.Itoa(count[i])
		}
		res[key] = append(res[key], s)
	}
	for _, v := range res {
		output = append(output, v)
	}
	return output
}

func find(group []string, query string) bool {
	for i := 0; i < len(group); i++ {
		if query == group[i] {
			return true
		}
	}
	return false
}
