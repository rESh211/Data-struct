package linkedlist

type LinkedList[T any] struct {
}

type Node[T any] struct {
	data T
	next *Node[T]
}
