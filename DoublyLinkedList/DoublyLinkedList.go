package doublyLinkedList

type Node1[T any] struct {
	data T
	prev *Node1[T]
	next *Node1[T]
}

type DoublyLinkedList[T any] struct {
	Count int
	tail  *Node1[T]
	head  *Node1[T]
}

func NewLinkedList1[T any]() *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{}
}

// Метод добавление в начало.
func (doublyLinkedList *DoublyLinkedList[T]) AddingBeginning(data T) {
	newNode := &Node1[T]{
		data: data,
	}
	if doublyLinkedList.head == nil {
		doublyLinkedList.head = newNode
		doublyLinkedList.tail = newNode
	} else {
		newNode.next = doublyLinkedList.head
		doublyLinkedList.head.prev = newNode
		doublyLinkedList.head = newNode
	}
	doublyLinkedList.Count++
}

// Метод добавление в конец.
func (doublyLinkedList *DoublyLinkedList[T]) AddingAtEnd(data T) {
	newNode := &Node1[T]{
		data: data,
	}
	if doublyLinkedList.head == nil {
		doublyLinkedList.head = newNode
		doublyLinkedList.tail = newNode
	} else {
		one := doublyLinkedList.head
		for one.next != nil {
			one = one.next
		}
		newNode.prev = one
		one.next = newNode
		doublyLinkedList.tail = newNode
	}
	doublyLinkedList.Count++
}

// Метод удаления из начало.
func (doublyLinkedList *DoublyLinkedList[T]) RemovalBeginning() {
	if doublyLinkedList.head == nil {
		return
	}
	if doublyLinkedList.head.next == nil {
		doublyLinkedList.head = nil
		doublyLinkedList.tail = nil
	} else {
		doublyLinkedList.head = doublyLinkedList.head.next
		doublyLinkedList.head.prev = nil
	}
	doublyLinkedList.Count--
}

// Метод удаления из конца.
func (doublyLinkedList *DoublyLinkedList[T]) RemovalEnd() {
	if doublyLinkedList.tail == nil {
		return
	}
	if doublyLinkedList.tail.next == nil {
		doublyLinkedList.head = nil
		doublyLinkedList.tail = nil
	} else {
		doublyLinkedList.tail = doublyLinkedList.tail.next
		doublyLinkedList.tail.next = nil
	}
	doublyLinkedList.Count--
}

/* Метод узла по значению.
func (doublyLinkedList *DoublyLinkedList[T]) Search(data T) *Node1[T] {
	number := doublyLinkedList.head
	for number != nil {
		if {
			return number
		}
		number = number.next // Исправлено
	}
	return nil
}*/
