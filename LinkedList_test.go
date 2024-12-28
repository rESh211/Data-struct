package main

import (
	"testing"

	//github.com/rESh211/Data-struct.git"

	linkedlist "github.com/rESh211/Data-struct.git/LinkedList"
	"github.com/stretchr/testify/assert"
)

func TestNewLinkedList(t *testing.T) {
	assert := assert.New(t)
	assert.NotPanics(func() {
		linkedlist.NewLinkedList[int]()
	})
}

func TestPush(t *testing.T) {
	assert := assert.New(t)
	assert.NotPanics(func() {
		list := linkedlist.NewLinkedList[int]()
		list.Push(251)
	})
}

func TestGet(t *testing.T) {
	assert := assert.New(t)
	assert.NotPanics(func() {
		list := linkedlist.NewLinkedList[int]()
		list.Push(251)
		ferst := list.Get(0)
		assert.Equal(251, *ferst, "проверка первого элемента")
	})
}
