package searchengine

import "container/heap"

// feature 6 rearrange the search results such that results from the same website do not show up together

// Solution
// the optimal way to solve this question is to use the greedy technique that using the heap data structure.
// the condition for failing occurs when the frequency of a character exceeds (n+1)/2
// 1) we will find the frequency of the most frequent character in the input string
// 2) if the most frequent character in the string has the frequency maxFreq, for a valid rearrangement
// of the string, the remaining characters should be at least maxFreq-1. We will verify this case.
// We can use the map in Go. We will simply return the initial string if this test fails.
// 3) We will be using a max heap to store the frequencies of the characters. We can use the MaxHeap
// data structure in Go to simulate a max heap. We will add(char, freq) tuples to a heap.
// 4) next, we will remove the most frequent element. If that element is not the last element in the
// result, we add it to the result.
// 5) if it is the last element of the result, we pick the second most frequent element and add that to the
// result. We also add the most frequent element back to the heap.
// 6) when we pop form the heap and append it to the result, we need to add it back to the heap if the
// frequency is not 0.
// 7) return the result.

func reorganizeResults(initailOrder string) string {
	freqMap := make(map[rune]int)
	// step 1
	for _, c := range initailOrder {
		freq := 1
		if _, ok := freqMap[c]; ok {
			freq = freqMap[c] + 1
		}
		if freq > (len(initailOrder)+1)/2 {
			return initailOrder
		}
		freqMap[c] = freq
	}

	pq := &MaxHeap{}
	heap.Init(pq)
	for i, _ := range freqMap {
		heap.Push(pq, []int{int(i), freqMap[i]})
	}

	result := ""
	for !pq.Empty() {
		first := pq.Top()
		heap.Pop(pq)
		if len(result) == 0 || first[0] != int(result[len(result)-1]) {
			result = result + string(first[0])
			first[1]--
			if first[1] > 0 {
				heap.Push(pq, first)
			}
		} else {
			second := pq.Top()
			heap.Pop(pq)
			result = result + string(second[0])
			second[1]--
			if second[1] > 0 {
				heap.Push(pq, second)
			}
			heap.Push(pq, first)
		}
	}
	return result
}
