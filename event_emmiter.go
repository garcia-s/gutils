package utils

import (
	"fmt"
	"sync"
)

type eventListener[T interface{}] struct {
	once    bool
	handler func(T)
}

type EventEmmitter[T interface{}] struct {
	mu        sync.RWMutex
	listeners map[string][]eventListener[T]
}

func NewEventEmmiter[T interface{}]() EventEmmitter[T] {
	return EventEmmitter[T]{
		listeners: make(map[string][]eventListener[T]),
	}
}

func (e *EventEmmitter[T]) Listen(eventName string, listener func(T)) {
	e.mu.Lock()
	defer e.mu.Unlock()

	if _, ok := e.listeners[eventName]; !ok {
		e.listeners[eventName] = []eventListener[T]{}
	}

	e.listeners[eventName] = append(
		e.listeners[eventName],
		eventListener[T]{
			false,
			listener,
		},
	)
}

func (e *EventEmmitter[T]) Off(eventName string, listener func(T)) {
	e.mu.Lock()
	defer e.mu.Unlock()
	listeners, ok := e.listeners[eventName]
	if !ok {
		return
	}

	for i, l := range listeners {
		if fmt.Sprintf("%p", l.handler) == fmt.Sprintf("%p", listener) {
			e.listeners[eventName] = append(listeners[:i], listeners[i+1:]...)
		}
	}
}

func (e *EventEmmitter[T]) Emit(eventName string, data T) {
	listeners, ok := e.listeners[eventName]
	if !ok {
		return
	}

	for i := 0; i < len(listeners); i++ {
		go listeners[i].handler(data)
	}
}
