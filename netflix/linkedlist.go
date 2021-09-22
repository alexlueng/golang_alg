package netflix

import (
	"fmt"
	"math/rand"
)

type LinkedListNode struct {
	key              int
	data             int
	next             *LinkedListNode
	arbitrartPointer *LinkedListNode
}

func createdLinkedList(lst []int) *LinkedListNode {
	var head *LinkedListNode
	var tail *LinkedListNode
	for _, x := range lst {
		newNode := &LinkedListNode{data: x}
		if head == nil {
			head = newNode
		} else {
			tail.next = newNode
		}
		tail = newNode
	}
	return head
}

func insertAtHead(head *LinkedListNode, data int) *LinkedListNode {
	newNode := &LinkedListNode{data: data}
	newNode.next = head
	return newNode
}

func insertAtTail(head *LinkedListNode, data int) *LinkedListNode {
	newNode := &LinkedListNode{data: data}
	if head == nil {
		return newNode
	}
	temp := head
	for temp.next != nil {
		temp = temp.next
	}
	temp.next = newNode
	return head
}

func createRandomList(length int) *LinkedListNode {
	var listHead *LinkedListNode
	for i := 0; i < length; i++ {
		listHead = insertAtHead(listHead, rand.Intn(100))
	}
	return listHead
}

func toList(head *LinkedListNode) []int {
	var lst []int
	temp := head
	for temp.next != nil {
		lst = append(lst, temp.data)
		temp = temp.next
	}
	return lst
}

func display(head *LinkedListNode) {
	temp := head
	for temp.next != nil {
		fmt.Printf("%d", temp.data)
		temp = temp.next
		if temp != nil {
			fmt.Printf(", ")
		}
	}
	fmt.Println()
}

func mergeAlternating(list1, list2 *LinkedListNode) *LinkedListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	head := list1
	for list1.next != nil && list2 != nil {
		temp := list2
		list2 = list2.next

		temp.next = list1.next
		list1.next = temp
		list1 = temp.next
	}
	if list1.next == nil {
		list1.next = list2
	}
	return head
}

func isEqual(list1, list2 *LinkedListNode) bool {
	if list1 == list2 {
		return true
	}
	for list1 != nil && list2 != nil {
		if list1.data != list2.data {
			return false
		}
		list1 = list1.next
		list2 = list2.next
	}
	return (list1 == list2)
}

type DoubleLinkedListNode struct {
	key  int
	data int
	next *DoubleLinkedListNode
	prev *DoubleLinkedListNode
}

type DoubleLinkedList struct {
	head *DoubleLinkedListNode
	tail *DoubleLinkedListNode
	size int
}

func (d *DoubleLinkedList) InsertAtHead(key, data int) {
	newNode := &DoubleLinkedListNode{key: key, data: data}
	if d.head == nil {
		d.head = newNode
		d.tail = newNode
	} else {
		newNode.next = d.head
		d.head.prev = newNode
		d.head = newNode
	}
	d.size++
}

func (d *DoubleLinkedList) InsertAtTail(key, data int) {
	newNode := &DoubleLinkedListNode{key: key, data: data}
	if d.tail == nil {
		d.tail = newNode
		d.head = newNode
		newNode.next = nil
	} else {
		d.tail.next = newNode
		newNode.prev = d.tail
		d.tail = newNode
		newNode.next = nil
	}
	d.size++
}

func (d *DoubleLinkedList) GetHead() *DoubleLinkedListNode {
	return d.head
}

func (d *DoubleLinkedList) GetTail() *DoubleLinkedListNode {
	return d.tail
}

func (d *DoubleLinkedList) RemoveNode(node *DoubleLinkedListNode) *DoubleLinkedListNode {
	if node == nil {
		return nil
	}
	if node.prev != nil {
		node.prev.next = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	}
	if node == d.head {
		d.head = d.head.next
	}
	if node == d.tail {
		d.tail = d.tail.prev
	}
	d.size--
	return node
}

func (d *DoubleLinkedList) Remove(data int) {
	i := d.GetHead()
	for i != nil {
		if i.data == data {
			d.RemoveNode(i)
		}
		i = i.next
	}
}

func (d *DoubleLinkedList) RemoveHead() *DoubleLinkedListNode {
	return d.RemoveNode(d.head)
}

func (d *DoubleLinkedList) RemoveTail() *DoubleLinkedListNode {
	return d.RemoveNode(d.tail)
}

type LFUNode struct {
	key  int
	val  int
	freq int
	next *LFUNode
	prev *LFUNode
}

type LFUList struct {
	head *LFUNode
	tail *LFUNode
}

func (l *LFUList) Append(node *LFUNode) {
	if l.head == nil {
		l.head = node
	} else {
		l.tail.next = node
		node.prev = l.tail
	}
	l.tail = node
}

func (l *LFUList) DeleteNode(node *LFUNode) {
	if node.prev != nil {
		node.prev.next = node.next
	} else {
		l.head = node.next
	}

	if node.next != nil {
		node.next.prev = node.prev
	} else {
		l.tail = node.prev
	}
	node.next = nil
	node.prev = nil
}
