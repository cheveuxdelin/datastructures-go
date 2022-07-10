package stack

import (
	"errors"
)

type node[T any] struct {
	value T
	next  *node[T]
}

var (
	ErrEmptyStack error = errors.New("empty stack")
)

type Stack[T any] struct {
	head *node[T]
	tail *node[T]
	size int
}

func (s *Stack[T]) Push(x T) {
	var newNode *node[T] = &node[T]{x, nil}
	if s.size == 0 {
		s.head = newNode
		s.tail = newNode
	} else {
		newNode.next = s.head
		s.head = newNode
	}
	s.size += 1
}

func (s *Stack[T]) Pop() (T, error) {
	if s.size > 0 {
		var aux node[T] = *s.head
		s.head = s.head.next
		s.size--
		return aux.value, nil
	} else {
		var t T
		return t, ErrEmptyStack
	}
}

func (s *Stack[T]) Front() (T, error) {
	if s.size > 0 {
		return s.head.value, nil
	} else {
		var t T
		return t, ErrEmptyStack
	}
}

func (s *Stack[T]) Size() int {
	return s.size
}
