package datastructure

import (
	"fmt"
)

type Stack[T any] struct {
	slice []T
	size  int
}

// pushes the value into the stack
func (s *Stack[T]) Push(v T) {
	s.slice = append(s.slice, v)
	s.size++
}

// pops the value from the top of the stack and returns it. if the
// stack is empty it returns the default value and false
func (s *Stack[T]) Pop() (T, bool) {
	if s.Empty() {
		var def T
		return def, false
	}

	v := s.slice[s.size-1]
	s.slice = s.slice[:s.size-1]
	s.size--

	return v, true
}

// The top value of the stack i.e. the value of the last
func (s *Stack[T]) Top() T {
	if s.Empty() {
		var def T
		return def
	}

	return s.slice[s.size-1]
}

// Returns true if the stack is empty
func (s *Stack[T]) Empty() bool {
	return s.size == 0
}

// Returns the size of the stack
func (s *Stack[T]) Size() int {
	return s.size
}

func (s *Stack[T]) Values() []T {
	v := make([]T, s.Size())

	copy(v, s.slice)

	return v
}

func (s *Stack[T]) String() string {
	return fmt.Sprint(s.slice)
}
