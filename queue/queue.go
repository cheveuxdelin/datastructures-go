package queue

import (
	"errors"
)

type node[T any] struct {
	value T
	next  *node[T]
}

var (
	ErrEmptyQueue error = errors.New("empty queue")
)

type Queue[T any] struct {
	head *node[T]
	tail *node[T]
	size int
}

func (s *Queue[T]) Push(x T) {
	var newNode *node[T] = &node[T]{x, nil}
	if s.size == 0 {
		s.head = newNode
		s.tail = newNode
	} else {
		s.tail.next = newNode
		s.tail = newNode
	}
	s.size += 1
}

func (s *Queue[T]) Pop() (T, error) {
	if s.size > 0 {
		var aux node[T] = *s.head
		s.head = s.head.next
		s.size--
		return aux.value, nil
	} else {
		var t T
		return t, ErrEmptyQueue
	}
}

func (s *Queue[T]) Front() (T, error) {
	if s.size > 0 {
		return s.head.value, nil
	} else {
		var t T
		return t, ErrEmptyQueue
	}
}

func (s *Queue[T]) Size() int {
	return s.size
}
