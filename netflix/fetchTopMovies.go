package netflix

// feature 2 Enable the user to view the top-rated movies worldwide, given that we have movie rankings available separately for different geographic regions.

// this question equals merge k linklists
// Solution
// since our task involves multiple lists, we should divide the problem into multiple tasks.

func merge2SortedLists(l1, l2 *LinkedListNode) *LinkedListNode {
	dummy := &LinkedListNode{data: -1}
	prev := dummy

	for l1 != nil && l2 != nil {
		if l1.data <= l2.data {
			prev.next = l1
			l1 = l1.next
		} else {
			prev.next = l2
			l2 = l2.next
		}
		prev = prev.next
	}
	if l1 == nil {
		prev.next = l2
	} else {
		prev.next = l1
	}
	return dummy.next
}

func mergeKSortedLists(lists []*LinkedListNode) *LinkedListNode {

	if len(lists) == 0 {
		return &LinkedListNode{data: -1}
	}
	res := lists[0]
	for i := 1; i < len(lists); i++ {
		res = merge2SortedLists(res, lists[i])
	}
	return res

}
