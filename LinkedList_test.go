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