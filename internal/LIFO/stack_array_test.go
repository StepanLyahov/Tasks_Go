package LIFO

import (
	"log"
	"testing"
)

func TestNewArrayStackIntWithNegativeSize(t *testing.T) {
	var _, err = NewArrayStackInt(-1)

	if err == nil {
		t.Fatalf("Err must be exist")
	}
}

func TestNewArrayStackInt(t *testing.T) {
	_, err := NewArrayStackInt(10)

	if err != nil {
		t.Fatalf("Err must be nil")
	}
}

func TestArrayStackInt_Push(t *testing.T) {
	stack, err := NewArrayStackInt(10)

	if err != nil {
		t.Fatalf("Err must be nil. Err: %v", err)
	}

	err = stack.Push(1)

	if err != nil {
		t.Fatalf("Pushing error. Err: %v", err)
	}
}

func TestArrayStackInt_Pop(t *testing.T) {
	stack, err := NewArrayStackInt(1)

	if err != nil {
		t.Fatalf("Err must be nil. Err: %v", err)
	}

	val := 1

	err = stack.Push(val)
	err = stack.Push(val)
	_, err = stack.Pop()
	pop, err := stack.Pop()
	if err != nil {
		t.Fatalf("Poping error. Err: %v", err)
	}

	if pop != val {
		t.Fatalf("Value in push and pop is not equel. Pop %v Val %v", pop, val)
	}
}

func TestArrayStackInt_Pop_StackIsVoid(t *testing.T) {
	stack, err := NewArrayStackInt(10)

	if err != nil {
		t.Fatalf("Err must be nil. Err: %v", err)
	}

	_, err = stack.Pop()
	if err == nil {
		t.Fatalf("Err must be exist")
	}
	log.Print(err)
}

func TestArrayStackInt_Get(t *testing.T) {
	stack, err := NewArrayStackInt(10)

	if err != nil {
		t.Fatalf("Err must be nil. Err: %v", err)
	}

	val := 1
	err = stack.Push(val)
	get, err := stack.Get()
	if err != nil {
		t.Fatalf("Getting error. Err: %v", err)
	}

	if get != val {
		t.Fatalf("Value in push and get is not equel. Pop %v Val %v", get, val)
	}
}

func TestArrayStackInt_Get_StackIsVoid(t *testing.T) {
	stack, err := NewArrayStackInt(10)

	if err != nil {
		t.Fatalf("Err must be nil. Err: %v", err)
	}

	_, err = stack.Get()
	if err == nil {
		t.Fatalf("Err must be exist")
	}
	log.Print(err)
}
