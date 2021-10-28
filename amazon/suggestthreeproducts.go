package amazon

import "sort"

// feature 2 some customers have won a coupon with $200 to buy up to three items. Suggest a bundle of
// three items to these lucky users.
// Solution
// this feature can also descripe as three sum problem
// 1) check if any items are greater than 200
// 2) check if the current value and previous value are the same. if this is true, then we will skip the
// current element
// 3) call the twoProducts helper function with the itemPrices list and the current element's index,i

func twoProducts(itemPrices []int, i int, res [][]int) [][]int {
	seen := make(map[int]bool)
	j := i + 1

	for j < len(itemPrices) {
		complement := 200 - itemPrices[i] - itemPrices[j]
		if _, ok := seen[complement]; ok {
			res = append(res, []int{itemPrices[i], itemPrices[j], complement})
			for j+1 < len(itemPrices) && itemPrices[j] == itemPrices[j+1] {
				j++
			}
		}
		seen[itemPrices[j]] = true
		j++
	}
	return res
}

func suggestThreeProducts(itemPrices []int) [][]int {
	var res [][]int
	sort.Ints(itemPrices)
	for i := 0; i < len(itemPrices); i++ {
		price := itemPrices[i]
		if price > 200 {
			break
		}
		if i == 0 || itemPrices[i-1] != price {
			res = twoProducts(itemPrices, i, res)
		}
	}
	return res
}
