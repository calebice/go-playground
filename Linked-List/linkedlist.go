package linkedlist

import (
	"fmt"
	"strings"
)

// Item a data holding representation
type Item interface{}

// Node a data holding pointer
type Node struct {
	data Item
	next *Node
}

// LinkedList implements a singlely linked list using interface type
type LinkedList struct {
	head *Node
	size int
}

// Append places the input value at the end of the linked list. Returns item inserted
func (ll *LinkedList) Append(data Item) {
	insertVal := &Node{data: data}
	if ll.head == nil {
		ll.head = insertVal
		ll.size++
		return
	}

	temp := ll.head
	for temp.next != nil {
		temp = temp.next
	}
	ll.size++

	temp.next = insertVal
	return
}

// Insert attempts to insert data at the provided location. Returns error for invalid indices
func (ll *LinkedList) Insert(i int, data Item) (Item, error) {
	temp := ll.head
	node := &Node{data: data, next: nil}
	if i < 0 || i > ll.size {
		return nil, fmt.Errorf("Invalid index provided: %v", i)
	}

	if i == 0 {
		node.next = ll.head
		ll.head = node
		ll.size++
		return data, nil
	}

	for j := 0; j < i-1; j++ {
		temp = temp.next
	}
	node.next = temp.next
	temp.next = node
	ll.size++
	return data, nil
}

// RemoveAt removes the provided index returns error for invalid indices
func (ll *LinkedList) RemoveAt(i int) (Item, error) {
	temp := ll.head
	if i < 0 || i > ll.size {
		return nil, fmt.Errorf("Invalid index provided: %v", i)
	}

	if i == 0 {
		removeVal := temp.data
		temp = temp.next
		ll.size--
		return removeVal, nil
	}

	for j := 0; j < i-1; j++ {
		temp = temp.next
	}
	removeVal := temp.next.data
	temp.next = temp.next.next
	ll.size--
	return removeVal, nil
}

// IndexOf finds the index of the provided value or -1 if it does not contain value
func (ll *LinkedList) IndexOf(data Item) int {
	temp := ll.head
	index := 0
	for temp.next != nil {
		if temp.data == data {
			return index
		}
		temp = temp.next
		index++
	}
	if temp.data == data {
		return index
	}
	return -1
}

// Size returns size of linkedlist
func (ll LinkedList) Size() int {
	return ll.size
}

// Head returns head of the linked list
func (ll LinkedList) Head() *Node {
	return ll.head
}

func (ll LinkedList) String() string {
	var stringRep string
	temp := ll.head
	stringRep = fmt.Sprintf("head[%v] -> ", temp.data)
	for temp.next != nil {
		temp = temp.next
		stringRep = fmt.Sprintf("%v[%v] -> ", stringRep, temp.data)
	}
	stringRep = strings.TrimSuffix(stringRep, " -> ")
	return stringRep
}
