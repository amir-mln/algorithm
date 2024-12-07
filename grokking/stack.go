package grokking

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

/*
Removes the top most element of the stack and returns it.
returns default value of T and false when the
stack is empty.
*/
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

/*
Removes the bottom most element of the stack and returns it.
returns default value of T and false when the
stack is empty.
*/
func (s *Stack[T]) PopBottom() (T, bool) {
	if s.Empty() {
		var def T
		return def, false
	}

	v := s.slice[0]
	s.slice = s.slice[1:]
	s.size--

	return v, true
}

/*
The top most element of the stack.
returns default value of T and false when the
stack is empty.
*/
func (s *Stack[T]) Top() (T, bool) {
	if s.Empty() {
		var def T
		return def, false
	}

	return s.slice[s.size-1], true
}

/*
The bottom most element of the stack.
returns default value of T and false when the
stack is empty.
*/
func (s *Stack[T]) Bottom() (T, bool) {
	if s.Empty() {
		var def T
		return def, false
	}

	return s.slice[0], true
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
