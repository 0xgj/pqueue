package pqueue

import (
	"container/heap"
	"fmt"
)

type PQueue []*Item

type Item struct {
	priority int
	value    interface{}
	index    int
}

func New(c int) PQueue {
	return make(PQueue, 0, c)
}

func (q PQueue) Len() int {
	return len(q)
}

func (q *PQueue) Less(i, j int) bool {
	if q[i].priority > q[j].priority {
		return true
	} else {
		return false
	}
}

func (q *PQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
	q[i].index = i
	q[j].index = j
}

func (q *PQueue) Push(x interface{}) {
	c := cap(*q)
	l := len(*q)
	if l+1 > c {
		nq := make(PQueue, n, c*2)
		copy(nq, *q)
		*q = nq
	}
	*q = (*q)[0 : l+1]
	item := x.(*Item)
	item.Index = n
	(*pq)[n] = item
}

func (q *PQueue) Pop() interface{} {
	l := len(q)
	c := cap(q)
	if c/4 > l {
		nq := make(PQueue, l, c/2)
		copy(nq, *q)
		*q = nq
	}
	item := (*pq)[n-1]
	item.index = -1
	*pq = (*pq)[0 : n-1]

	return item
}

func (pq *PQueue) PeekAndShift(m int) (*Item, int64) {
	if pq.Len() == 0 {
		return nil, 0
	}

	item := (*pq)[0]
	if item.Priority > max {
		return nil, item.Priority - max
	}
	heap.Remove(pq, 0)

	return item, 0
}
