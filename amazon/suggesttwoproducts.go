package amazon

// feature 1 All orders above a certain total amount are eligible for free delivery. If a customer's
// total bill is near the free-shipping threshold, recommend items to add to cart to avail free delivery
// Solution
// This is a two sum question

func suggestTwoProducts(itemPrices []int, amount int) (int, int) {
	buffDict := make(map[int]int)
	for i := 0; i < len(itemPrices); i++ {
		price := itemPrices[i]
		remaining := amount - price
		if _, ok := buffDict[remaining]; !ok {
			buffDict[price] = i
		} else {
			return buffDict[remaining], i
		}
	}
	return 0, 0
}
