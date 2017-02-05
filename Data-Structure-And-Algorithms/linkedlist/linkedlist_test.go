package linkedlist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeletion0Items(t *testing.T) {
	//list with one item
	list := LinkedList{}
	list.DeleteNode(4)
	assert.Equal(t, list.Size(), 0, "size should be zero")
}

func TestInsertDelete1Item(t *testing.T) {
	//list with one item
	list := LinkedList{}
	list.InsertList(1)
	assert.Equal(t, list.GetTailValue(), 1, "tail should be 1")
	assert.Equal(t, list.GetHeadValue(), 1, "head should be 1")
	list.DeleteNode(1)
	assert.Equal(t, list.Size(), 0, "size should be zero")
	assert.Equal(t, list.GetTailValue(), nil, "tail should be 1")
	assert.Equal(t, list.GetHeadValue(), nil, "tail should be 1")
}

func TestInsertDelete2Item1(t *testing.T) {
	//list with one item
	list := LinkedList{}
	list.InsertList(1)
	list.InsertList(2)
	assert.Equal(t, 1, list.GetTailValue(), "tail should be 1")
	assert.Equal(t, 2, list.GetHeadValue(), "head should be 2")
	list.DeleteNode(2)
	assert.Equal(t, 1, list.Size(), "size should be one")
	assert.Equal(t, 1, list.GetTailValue(), "tail should be 1")
	assert.Equal(t, 1, list.GetHeadValue(), "head should be 1")
	list.InsertList(2)
	list.DeleteNode(1)
	assert.Equal(t, 1, list.Size(), "size should be one")
	assert.Equal(t, 2, list.GetTailValue(), "tail should be 1")
	assert.Equal(t, 2, list.GetHeadValue(), "head should be 1")
}

func TestInsertListEnd(t *testing.T) {
	list := LinkedList{}
	list.InsertListEnd(1)
	assert.Equal(t, 1, list.GetTailValue(), "tail should be 1")
	assert.Equal(t, 1, list.GetHeadValue(), "head should be 1")
	list.InsertListEnd(2)
	assert.Equal(t, 2, list.GetTailValue(), "tail should be 2")
	assert.Equal(t, 1, list.GetHeadValue(), "head should be 1")
	list.InsertListEnd(3)
	assert.Equal(t, 3, list.GetTailValue(), "tail should be 3")
	assert.Equal(t, 1, list.GetHeadValue(), "head should be 1")
	assert.Equal(t, 3, list.Size(), "size should be 3")
}
