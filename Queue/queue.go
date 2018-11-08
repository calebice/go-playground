package queue

import (
	"fmt"
	"strings"
)

// Item is a generic(ish) type for queues
type Item interface{}

// Queue is a list based queue representation
type Queue struct {
	items []Item
}

// New instantiates an array of items
func (s *Queue) New() *Queue {
	s.items = []Item{}
	return s
}

// Enqueue adds an element to the queue
func (s *Queue) Enqueue(item Item) {
	s.items = append(s.items, item)
}

// Dequeue pops an element off the queue
func (s *Queue) Dequeue() Item {

	if s.IsEmpty() {
		return nil
	}
	item := s.items[0]
	s.items = s.items[1:]
	return item
}

// Size returns the size of the array
func (s Queue) Size() int {
	return len(s.items)
}

// Front returns the front of the queue
func (s Queue) Front() Item {
	if s.IsEmpty() {
		return nil
	}
	return s.items[0]
}

// IsEmpty determines if the queue is empty
func (s Queue) IsEmpty() bool {
	return len(s.items) == 0
}

func (s Queue) String() string {
	var retString string
	for _, item := range s.items {
		retString = fmt.Sprintf("%v, %v", retString, item)
	}
	retString = strings.TrimPrefix(retString, ", ")
	return retString
}
