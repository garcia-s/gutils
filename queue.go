package gutils

import (
	"errors"
)

type Queue[T any] struct {
	head   *qNode[T]
	tail   *qNode[T]
	length int
}

type qNode[T any] struct {
	data T
	next *qNode[T]
}

func NewQueue[T any]() Queue[T] {
	return Queue[T]{
		length: 0,
	}
}

func (q *Queue[T]) Enqueue(value T) {
	tail := &qNode[T]{
		data: value,
	}

	if q.head == nil {
		q.head = tail
	} else {
		q.tail.next = tail
	}

	q.tail = tail
	q.length++
}

func (q *Queue[T]) Poll() (value *T, err error) {
	if q.length == 0 {
		return nil, errors.New("Trying to poll an empty Queue")
	}

	current := q.head
	q.head = current.next
	q.length--
	return &current.data, nil
}

func (q *Queue[T]) Peek() (value *T, err error) {
	if q.length == 0 {
		return nil, errors.New("There are no items in this queue")
	
    }

	current := *&q.head.data
	return &current, nil
}

func (q Queue[T]) Length() int {
	return q.length
}
