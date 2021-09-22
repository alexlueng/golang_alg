package netflix

import "fmt"

// feature 5 for the client application, we want to implement a cache with a replacement strategy that replace the least recently watched title

// what we want is that a user can view the n titles that they've watched or accessed most recently.
// we need a data structure: 1) should maintain titles in order of time since last access
// 2) if the data structure is at its capacity, an insertion should replace the least recently accessed item.

// Solution
// double linked list would be an ideal way.
// 1) if the element exists in the hashmap, our first step is to move the accessed element to the tail of the linked list
// 2) if the data structure is at its capacity, remove the element at the head of the linked list and the hashmap. Then, we add the new element at the tail of the linked list and in the hash map
// 3) retrieve the element and return

type LRUCache struct {
	capacity  int
	cache     map[int]*DoubleLinkedListNode
	cacheVals DoubleLinkedList
}

func (lr *LRUCache) Get(key int) *DoubleLinkedListNode {
	if _, ok := lr.cache[key]; !ok {
		return nil
	}
	value := lr.cache[key].data
	lr.cacheVals.RemoveNode(lr.cache[key])
	lr.cacheVals.InsertAtTail(key, value)
	return lr.cacheVals.GetTail()
}

func (lr *LRUCache) Set(key, value int) {
	if _, ok := lr.cache[key]; !ok {
		if lr.cacheVals.size >= lr.capacity {
			lr.cacheVals.InsertAtTail(key, value)
			lr.cache[key] = lr.cacheVals.GetTail()
			delete(lr.cache, lr.cacheVals.GetHead().key)
			lr.cacheVals.RemoveHead()
		} else {
			lr.cacheVals.InsertAtTail(key, value)
			lr.cache[key] = lr.cacheVals.GetTail()
		}
	} else {
		lr.cacheVals.RemoveNode(lr.cache[key])
		lr.cacheVals.InsertAtTail(key, value)
		lr.cache[key] = lr.cacheVals.GetTail()
	}
}

func (lr *LRUCache) Print() string {
	result := ""
	curr := lr.cacheVals.head
	for curr != nil {
		res := fmt.Sprintf("(%v, %v)", curr.key, curr.data)
		result += res
		curr = curr.next
	}
	// fmt.Println()
	return result
}
