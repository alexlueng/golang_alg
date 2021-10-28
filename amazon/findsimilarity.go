package amazon

import "reflect"

// feature 6 show product recommendations with items that are frequently viewed together
// Solution
// use a sliding window on the products dataset and keep a hashtable referring to the occurrences of the
// products in the sliding window.

func findSimilarity(products []int, candidates []int) []int {
	prodN := len(products)
	candN := len(candidates)

	var output []int
	if prodN < candN {
		return output
	}

	candCount := make(map[int]int)
	prodCount := make(map[int]int)

	for _, i := range candidates {
		if _, ok := candCount[i]; ok {
			candCount[i] += 1
		} else {
			candCount[i] = 1
		}
	}

	for i := 0; i < prodN; i++ {
		k := products[i]
		if _, ok := prodCount[k]; ok {
			prodCount[k] += 1
		} else {
			prodCount[k] = 1
		}

		if i >= candN {
			k = products[i-candN]
			if prodCount[k] == 1 {
				delete(prodCount, k)
			} else {
				prodCount[k] -= 1
			}
		}

		if reflect.DeepEqual(candCount, prodCount) {
			output = append(output, i-candN+1)
		}
	}
	return output
}
