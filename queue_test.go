package gutils
import (
	"testing"
)

func TestEnqueueingWithEmptyList(t *testing.T) {
	queue := NewQueue[int]()
	queue.Enqueue(33)
	if queue.Length() == 0 {
		t.Fatalf("Enqueing failed to update the Length")
	}
}

func TestPollingEmptyQueue(t *testing.T) {
	queue := NewQueue[int]()
	_, err := queue.Poll()
	if err == nil {
		t.Fatalf("Queue not returning error being polled while empty")
	}

}

func TestPollWithSingleItem(t *testing.T) {
	queue := NewQueue[string]()
	queue.Enqueue("Hello")
	item, _ := queue.Poll()

	if item == nil {
		t.Fatalf("The Poll item returns a nil pointer")
	}
	if *item != "Hello" {
		t.Fatalf("Retrieved data is not working as spected")
	}
}

func TestPeekWithEmptyQueue(t *testing.T) {
	queue := NewQueue[string]()
	_, err := queue.Peek()
	if err == nil {
		t.Fatal("Peeking with an empty queue doesn't return an error")
	}
}

func TestPeekDoesntMutateQueueItem(t *testing.T) {
	queue := NewQueue[int]()
	queue.Enqueue(33)
	peeked, _ := queue.Peek()
	polled, _ := queue.Poll()

	if *peeked != 33 {
		t.Fatal("The peeked value is not the same as the enqueued value")
	}

	if *polled != 33 {
		t.Fatal("The polled value is not the same as the enqueued value")
	}
	//Pointers must be different
	if polled == peeked {
		t.Fatal("Changing a peeked element modifies the item inside the queue")
	}

	if queue.Length() != 0 {
		t.Fatal("The poll function is not decreasing the Lenght of the Queue")
	}
}

func TestQueueWithStructs(t *testing.T) {
	type Stub struct {
		name string
	}
	queue := NewQueue[Stub]()
	queue.Enqueue(Stub{name: "hello"})
	polled, _:= queue.Poll()

	if polled == nil {
		t.Fatal("Polled pointer is nil")
	}

	if polled.name != "hello" {
		t.Fatal("The polled struct values are not the same as the enqueued one")
	}

}
