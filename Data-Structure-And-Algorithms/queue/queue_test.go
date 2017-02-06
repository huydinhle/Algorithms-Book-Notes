package queue

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueue(t *testing.T) {
	q := Queue{}
	fmt.Printf("q.getList().Size() = %+v\n", q.GetList().Size())
	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)

	assert.Equal(t, 3, q.GetList().Size(), "size should be 3")
	assert.Equal(t, 1, q.DeQueue(), "should be 1")
	assert.Equal(t, 2, q.DeQueue(), "should be 2")
	assert.Equal(t, 3, q.DeQueue(), "should be 3")
}
