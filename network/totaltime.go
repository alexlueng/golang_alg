package network

import (
	"container/list"
)

// feature 1 a message is transmitted from the server to the clients. Given the delays for each hop
// in the network, we want to determine the earliest time at which the message will be received by
// all the clients, assuming there are no errors along the way.
// Solution(Spanning tree protocol)
// Eventually, we want to find the maximum value over all root-to-leaf paths.
// 1) build the adjacency list, where each server node points to its children in the spanning tree
// 2) create a tuple containing the main-server id and its delay time as(id, time), and insert it into a queue
// 3) traverse over the nodes while summing the maximum value at each level. Keep inserting each node and
// the sum until the node we just inserted into the queue.
// 4) return computed sum

type Pair struct {
	first, second int
}

func totalTime(mainServerId int, parents, delays []int) int {
	n := len(parents)
	if n <= 1 {
		return 0
	}

	res := 0
	children := make(map[int][]int)

	var numbers []int
	for i := 0; i < len(parents); i++ {
		numbers = append(numbers, parents[i])
	}

	for index := 0; index < len(numbers); index++ {
		val := numbers[index]
		var tempList []int
		if _, ok := children[val]; ok {
			tempList = children[val]
		}
		tempList = append(tempList, index)
		children[val] = tempList
	}

	queue := list.New()
	queue.PushBack(Pair{mainServerId, delays[mainServerId]})

	for queue.Len() != 0 {
		node := queue.Front()
		queue.Remove(node)
		curr := node.Value.(Pair)
		currId := curr.first
		currTime := curr.second

		// calculate max
		if res < currTime {
			res = currTime
		}
		var temp []int
		if _, ok := children[currId]; ok {
			temp = children[currId]
		}
		for _, child := range temp {
			queue.PushBack(Pair{child, currTime + delays[child]})
		}
	}
	return res
}
