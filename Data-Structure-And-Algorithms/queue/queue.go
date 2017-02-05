package queue

import (
	"github.com/huydinhle/Algorithms-Book-Notes/Data-Structure-And-Algorithms/linkedlist"
)

// Queue type
type Queue struct {
	list *linkedlist.LinkedList
}

func (q Queue) EnQueue(element interface{}) {
	q.list.InsertListEnd(element)
}

func (q Queue) DeQueue() interface{} {
	if q.list.Size() == 0 {
		return nil
	}
	element := q.list.DeleteNode(q.list.GetHeadValue())
	return element
}

func (q Queue) GetList() *linkedlist.LinkedList {
	return q.list
}
