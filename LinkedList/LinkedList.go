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

func (linkedList *LinkedList[T]) Push(value T) {
	if linkedList.Node == nil {
		linkedList.Node = &Node[T]{data: value}
		return
	}
	next := linkedList.Node
	for next != nil && next.next != nil {
		next = next.next
	}
	next.next = &Node[T]{data: value}
}
