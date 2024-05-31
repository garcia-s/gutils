package utils

import (
	"sync"
	"testing"
)

type EvArgs struct {
	Mess string
	Num  int
}

func TestListen(t *testing.T) {
	ev := NewEventEmmiter[int]()

	wg := sync.WaitGroup{}
	wg.Add(1)

	ev.Listen("test", func(int) {
		defer wg.Done()
	})

	ev.Emit("test", 3)
	wg.Wait()
}

func TestOff(t *testing.T) {
	ev := NewEventEmmiter[EvArgs]()

	listener := func(EvArgs) {

	}

	for i := 0; i < 30; i++ {
		ev.Listen("test", listener)
	}

	for i := 0; i < 30; i++ {
		ev.Off("test", listener)
	}

	listeners, ok := ev.listeners["test"]

	if !ok {
		t.Errorf("There is no listener's list in the emmitter despite adding 30 events")
	}

	if len(listeners) != 0 {
		t.Errorf("Pointers to the same handlerFunction will result in the deletion of all the functions")
	}
}
