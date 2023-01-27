package priority

import "golang.org/x/exp/constraints"

type fibHeap[T any] struct {
	count uint
	root  *fibNode[T]
	min   *fibNode[T]
	pred  predicate
}

// newFibHeap TODO
func newFibHeap[T any](p predicate[T], items ...T) *fibHeap[T] {
	return &fibHeap[T]{}
}

// MinHeapFib TODO
func MinHeapFib[T constraints.Ordered](items ...T) Queue[T] {
	return newFibHeap(
		func(a, b T) bool { return a < b },
		items...,
	)
}

// MaxHeapFib TODO
func MaxHeapFib[T constraints.Ordered](items ...T) Queue[T] {
	return newFibHeap(
		func(a, b T) bool { return a > b },
		items...,
	)
}

// IsEmpty returns true if the fibonacci heap is empty - false otherwise
func (fh *fibHeap[T]) IsEmpty() bool {
	return fh.root == nil
}

// Peek provides the highest priority item without removing it from the fibonacci heap
// boolean retured indicates if the operation was successful i.e returns false if structure is empty
func (fh *fibHeap[T]) Peek() (top T, ok bool) {
	// TODO
	return
}

// Pop provides the highest priority item and removes it from the fibonacci heap
// boolean retured indicates if the operation was successful i.e returns false if structure is empty
func (fh *fibHeap[T]) Pop() (top T, ok bool) {
	// TODO
	return
}

// Push adds an item to the fibonacci heap
func (fh *fibHeap[T]) Push(item T) {

	n := &fibNode[T]{value: item}
	n.left, n.right = n, n

	if fh.root == nil { // item is the first item in the heap
		fh.root = n
	} else { // add item to the root list
		n.left = fh.root.left
		n.right = fh.root
		fh.root.left.right = n
		fh.root.left = n
	}

	if fh.min == nil || fh.pred(n, fh.min.value) {
		fh.min = n
	}
}

// fibNode represents a node in a fibonnaci heap
type fibNode[T any] struct {
	value  T
	parent *fibNode[T]
	child  *fibNode[T]
	left   *fibNode[T]
	right  *fibNode[T]
	degree uint // numbe of children
	marked bool // indicates if node has lost a child since the last time it was made the child of another node
}
