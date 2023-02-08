package priority

import "golang.org/x/exp/constraints"

// binaryHeap provides a binary heap
type binaryHeap[T comparable] struct {
	nodes []T
	pred  predicate[T]
	index map[T]int
}

// NewBinaryHeap returns a binary-heap based priority queue
func NewBinaryHeap[T comparable](p predicate[T], items ...T) *binaryHeap[T] {
	bh := &binaryHeap[T]{
		pred:  p,
		index: make(map[T]int),
	}

	for _, item := range items {
		bh.Push(item)
	}

	return bh
}

// MinHeapBinary TODO
func MinHeapBinary[T constraints.Ordered](items ...T) *binaryHeap[T] {
	return NewBinaryHeap(lessThan[T], items...)
}

// MaxHeapBinary TODO
func MaxHeapBinary[T constraints.Ordered](items ...T) *binaryHeap[T] {
	return NewBinaryHeap(greaterThan[T], items...)
}

// IsEmpty returns true if the binary heap is empty - false otherwise
func (bh *binaryHeap[T]) IsEmpty() bool {
	return len(bh.nodes) == 0
}

// Peek provides the highest priority item without removing it from the binary heap
// boolean retured indicates if the operation was successful i.e returns false if structure is empty
func (bh *binaryHeap[T]) Peek() (top T, ok bool) {
	if bh.IsEmpty() {
		return
	}

	return bh.nodes[0], true
}

// Pop provides the highest priority item and removes it from the binary heap
// boolean retured indicates if the operation was successful i.e returns false if structure is empty
func (bh *binaryHeap[T]) Pop() (top T, ok bool) {
	if bh.IsEmpty() {
		return
	}

	top = bh.nodes[0]
	delete(bh.index, top)

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
func (bh *binaryHeap[T]) Push(item T) {
	index := len(bh.nodes)
	bh.nodes = append(bh.nodes, item)
	bh.index[item] = index
	bh.upHeapify(len(bh.nodes) - 1)
}

func (bh *binaryHeap[T]) Update(t T, update UpdateFunc[T]) {
	ix, ok := bh.index[t]
	if !ok { // item is not in the heap
		return
	}

	bh.nodes[ix] = update(bh.nodes[ix])
	bh.upHeapify(ix)
}

// upHeapify fixes the heap structure bottom-up recursively
func (bh *binaryHeap[T]) upHeapify(child int) {
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
func (bh *binaryHeap[T]) downHeapify(parent int) {

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
func (bh *binaryHeap[T]) swap(i, j int) {
	ns := bh.nodes
	ns[i], ns[j] = ns[j], ns[i]
	bh.index[ns[i]] = i
	bh.index[ns[j]] = j
}

// checkPriority of items at given indices
func (bh *binaryHeap[T]) checkPriority(i, j int) bool {
	return bh.pred(bh.nodes[i], bh.nodes[j])
}
