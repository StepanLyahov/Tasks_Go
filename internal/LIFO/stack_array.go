package LIFO

import "errors"

type ArrayStackInt struct {
	buffer  []int
	pointer int
}

func NewArrayStackInt(initial int) (StackInt, error) {
	if initial < 1 {
		return &ArrayStackInt{}, errors.New("size must be more 0")
	}

	buffer := make([]int, initial)

	stack := ArrayStackInt{
		buffer:  buffer,
		pointer: -1,
	}

	return &stack, nil
}

func (a *ArrayStackInt) Push(i int) error {
	if a.pointer == len(a.buffer)-1 {
		a.pointer++
		a.buffer = append(a.buffer, i)
		return nil
	}

	a.pointer++
	a.buffer[a.pointer] = i
	return nil
}

func (a *ArrayStackInt) Pop() (int, error) {
	if a.pointer == -1 {
		return 0, errors.New("stack is void")
	}

	val := a.buffer[a.pointer]
	a.pointer--

	return val, nil
}

func (a *ArrayStackInt) Get() (int, error) {
	if a.pointer == -1 {
		return 0, errors.New("stack is void")
	}

	return a.buffer[a.pointer], nil
}
