package lib

type IntSet struct {
	Data map[int]struct{}
}

func (s *IntSet) Add(item int) {
	if !s.Contains(item) {
		s.Data[item] = struct{}{}
	}
}

func (s *IntSet) Contains(item int) bool {
	_, ok := s.Data[item]
	return ok
}

func (s *IntSet) Remove(item int) {
	if s.Contains(item) {
		delete(s.Data, item)
	}
}

func (s *IntSet) Len() int {
	return len(s.Data)
}

type PtrSet[V any] struct {
	Data map[*V]struct{}
}

func MakePtrSet[V any]() (s PtrSet[V]) {
	s.Data = make(map[*V]struct{})
	return
}

func (s *PtrSet[V]) Add(item *V) {
	if !s.Contains(item) {
		s.Data[item] = struct{}{}
	}
}

func (s *PtrSet[V]) Contains(item *V) bool {
	_, ok := s.Data[item]
	return ok
}

func (s *PtrSet[V]) Remove(item *V) {
	if s.Contains(item) {
		delete(s.Data, item)
	}
}

func (s *PtrSet[V]) Len() int {
	return len(s.Data)
}
