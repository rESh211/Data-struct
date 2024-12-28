package linkedlist

type LinkedList[T any] struct {
	Node *Node[T]
}

type Node[T any] struct {
	data T
	next *Node[T]
}

func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}
