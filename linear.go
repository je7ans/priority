package priority

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type linearPQ[T comparable] struct {
	items []T
	pred  predicate[T]
}

func NewLinearPriorityQueue[T comparable](p predicate[T], items ...T) *linearPQ[T] {
	q := &linearPQ[T]{pred: p}
	for _, t := range items {
		q.Push(t)
	}
	return q
}

func MinLinearQueue[T constraints.Ordered](items ...T) Queue[T] {
	return NewLinearPriorityQueue(lessThan[T], items...)
}

func MaxLinearQueue[T constraints.Ordered](items ...T) Queue[T] {
	return NewLinearPriorityQueue(greaterThan[T], items...)
}

func (q *linearPQ[T]) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *linearPQ[T]) Peek() (top T, ok bool) {
	var ix int
	if ix, ok = q.findHiPriority(); ok {
		top = q.items[ix]
	}

	return
}

func (q *linearPQ[T]) Pop() (top T, ok bool) {
	var ix int
	if ix, ok = q.findHiPriority(); ok {
		top = q.items[ix]
		q.items = append(q.items[:ix], q.items[ix+1:]...)
		fmt.Printf("popping %v - remaining: %v\n", top, q.items)
	}

	return
}

func (q *linearPQ[T]) Push(item T) {
	q.items = append(q.items, item)
}

func (q *linearPQ[T]) findHiPriority() (int, bool) {
	if q.IsEmpty() {
		return -1, false
	}

	top := 0
	for i := 1; i < len(q.items); i += 1 {
		if q.pred(q.items[i], q.items[top]) {
			top = i
		}
	}

	return top, true
}
