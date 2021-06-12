package LIFO

type StackInt interface {
	Push(int) error
	Pop() (int, error)
	Get() (int, error)
}
