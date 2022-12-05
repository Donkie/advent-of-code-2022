package lib

type Stack[V any] struct {
	data []V
}

// Returns the data at the specified index
func (s *Stack[V]) Get(idx int) *V {
	if idx >= s.Len() || idx < 0 {
		return nil
	}
	return &s.data[idx]
}

// Pushes a new element onto the stack
func (s *Stack[V]) Push(val V) {
	s.data = append(s.data, val)
}

// Pushes a list of new element onto the stack
// The elements are pushed so that they retain the same order when they're on the stack
// Just like dropping off a pile of plates onto an existing stack of plates
func (s *Stack[V]) PushN(values []V) {
	s.data = append(s.data, values...)
}

// Removes and returns the top-most element from the stack
func (s *Stack[V]) Pop() *V {
	li := s.PopN(1)
	if li == nil {
		return nil
	}
	return &(*li)[0]
}

// Removes and returns the top-most N elements from the stack
// The returned elements are retained in the order they were on the stack,
// just like picking up a sub-stack of a stack of plates
func (s *Stack[V]) PopN(amount int) *[]V {
	l := s.Len()
	if l < amount {
		return nil
	}

	output := s.data[l-amount : l]
	s.data = s.data[:l-amount]
	return &output
}

// Returns the top-most element of the stack
func (s *Stack[V]) Peek() *V {
	l := s.Len()
	if l == 0 {
		return nil
	}
	return &s.data[l-1]
}

// Returns the number of items in the stack
func (s *Stack[V]) Len() int {
	return len(s.data)
}
