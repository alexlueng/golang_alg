package amazon

import "math/rand"

// feature 3 up-sell a random product from a set of related products by recommending it to users at checkout

type UpsellProducts struct {
	productDict map[int]int
	productList []int
}

func (up *UpsellProducts) InsertProduct(prod int) bool {
	if _, ok := up.productDict[prod]; ok {
		return false
	}
	up.productDict[prod] = len(up.productList)
	up.productList = append(up.productList, prod)
	return true
}

func (up *UpsellProducts) RemoveProduct(prod int) bool {
	if _, ok := up.productDict[prod]; !ok {
		return false
	}
	last := up.productList[len(up.productList)-1]
	index := up.productDict[prod]
	up.productList[index] = last
	up.productDict[last] = index
	up.productList = up.productList[:len(up.productList)-1]
	delete(up.productDict, prod)
	return true
}

func (up *UpsellProducts) GetRandomProduct() int {
	return up.productList[rand.Intn(len(up.productList))]
}
