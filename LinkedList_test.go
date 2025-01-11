package main

import (
	"testing"

	//github.com/rESh211/Data-struct.git"

	linkedlist "github.com/rESh211/Data-struct.git/linkedlist"
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
		list.Push(52)
		second := list.Get(1)
		assert.Equal(52, *second, "проверка второго элемента")
		list.Push(487)
		third := list.Get(2)
		assert.Equal(487, *third, "проверка третьего элемента")
		second = list.Get(1)
		assert.Equal(52, *second, "повторная проверка второго элемента")
		assert.Nil(list.Get(3), "выход за границы индекса")
		assert.Panics(func() { list.Get(-1) }, "использование отрицательных значений")

	})
}

func TestSize(t *testing.T) {
	assert := assert.New(t)
	assert.NotPanics(func() {
		list := linkedlist.NewLinkedList[int]()
		assert.Equal(0, list.Size(), "проверка size пустого списка")
		list.Push(2)
		assert.Equal(1, list.Size(), " size не пустого списка")
		list.Push(3)
		assert.Equal(2, list.Size(), " size не пустого списка")
	})
}
