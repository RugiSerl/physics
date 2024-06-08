package math

type Tree[T any] struct {
	Data     T
	Children []Tree[T]
}

func (t *Tree[T]) Size() int {
	n := 1
	for i := 0; i < len(t.Children); i++ {
		n += t.Children[i].Size()
	}
	return n
}

func (t *Tree[T]) Height() int {
	n := -1
	for i := 0; i < len(t.Children); i++ {
		n = max(n, t.Children[i].Height())
	}
	return n
}
