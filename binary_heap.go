package priority

import "golang.org/x/exp/constraints"

// binHeap provides a binary heap
type binHeap[T any] struct {
	nodes []T
	pred  predicate[T]
}

// newBinaryHeap returns a binary-heap based priority queue
func newBinaryHeap[T any](p predicate[T], items ...T) *binHeap[T] {

	bh := &binHeap[T]{pred: p}

	for _, item := range items {
		bh.Push(item)
	}

	return bh
}

// MinHeapBinary TODO
func MinHeapBinary[T constraints.Ordered](items ...T) Queue[T] {
	return newBinaryHeap(
		func(a, b T) bool { return a < b },
		items...,
	)
}

// MaxHeapBinary TODO
func MaxHeapBinary[T constraints.Ordered](items ...T) Queue[T] {
	return newBinaryHeap(
		func(a, b T) bool { return a > b },
		items...,
	)
}

// IsEmpty returns true if the binary heap is empty - false otherwise
func (bh *binHeap[T]) IsEmpty() bool {
	return len(bh.nodes) == 0
}

// Peek provides the highest priority item without removing it from the binary heap
// boolean retured indicates if the operation was successful i.e returns false if structure is empty
func (bh *binHeap[T]) Peek() (top T, ok bool) {

	if bh.IsEmpty() {
		return
	}

	return bh.nodes[0], true
}

// Pop provides the highest priority item and removes it from the binary heap
// boolean retured indicates if the operation was successful i.e returns false if structure is empty
func (bh *binHeap[T]) Pop() (top T, ok bool) {

	if bh.IsEmpty() {
		return
	}

	top = bh.nodes[0]

	last := len(bh.nodes) - 1
	bh.swap(0, last)
	bh.nodes = bh.nodes[:last]

	if !bh.IsEmpty() {
		bh.downHeapify(0)
	}

	return top, true
}

// Push adds an item to the binary heap
// append the item to the end of the array (becomes the most bottom-right leaf node)
// fix the heap structure bottom-to-top
func (bh *binHeap[T]) Push(item T) {
	bh.nodes = append(bh.nodes, item)
	bh.upHeapify(len(bh.nodes) - 1)
}

// upHeapify fixes the heap structure bottoms-up recursively
func (bh *binHeap[T]) upHeapify(child int) {

	if child == 0 { // reached root node - no fixing required
		return
	}

	// if child has a higher priority than its parent node, swap with the parent and recurse
	if parent := (child - 1) / 2; bh.checkPriority(child, parent) {
		bh.swap(child, parent)
		bh.upHeapify(parent)
	}
}

// downHeapify fixes the heap structure recursively top-down starting at parent node
// if the parent node has lower priority than at least one of its children, swap with the
// highest priority child and recurse from that child node index
func (bh *binHeap[T]) downHeapify(parent int) {

	hiPriority := parent
	left := parent*2 + 1
	right := parent*2 + 2

	if left >= len(bh.nodes) { // parent is a leaf node
		return
	}

	if bh.checkPriority(left, hiPriority) { // left has higher priority than parent
		hiPriority = left
	}

	if right < len(bh.nodes) && bh.checkPriority(right, hiPriority) { // right has higher priority than both parent and left
		hiPriority = right
	}

	if parent != hiPriority { // fix required - swap & recurse
		bh.swap(parent, hiPriority)
		bh.downHeapify(hiPriority)
	}
}

// swap nodes at indices i and j
func (bh *binHeap[T]) swap(i, j int) {
	bh.nodes[i], bh.nodes[j] = bh.nodes[j], bh.nodes[i]
}

// checkPriority of items at given indices
func (bh *binHeap[T]) checkPriority(i, j int) bool {
	return bh.pred(bh.nodes[i], bh.nodes[j])
}
