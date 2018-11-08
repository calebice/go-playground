package set

// Item is a generic typed collection
type Item interface{}

// Set is an implementation of a set object
type Set struct {
	items map[Item]bool
}

// Add places an item into the set
func (s *Set) Add(t Item) *Set {
	if s.items == nil {
		s.items = make(map[Item]bool)
	}
	s.items[t] = true

	return s
}

// Remove removes an item from the set
func (s *Set) Remove(t Item) *Set {
	if s.items == nil {
		return nil
	}
	_, ok := s.items[t]
	if ok {
		delete(s.items, t)
	}
	return s
}

// Size calculates the size of the set
func (s *Set) Size() int {
	return len(s.items)
}

// Clear empties the set
func (s *Set) Clear() {
	s.items = make(map[Item]bool)
}

// Has checks to see if the set contains an item
func (s *Set) Has(t Item) bool {
	if s.items == nil {
		return false
	}
	return s.items[t]
}

// Items returns a list of all items in the set
func (s *Set) Items() []Item {
	items := []Item{}
	for item := range s.items {
		items = append(items, item)
	}

	return items
}
