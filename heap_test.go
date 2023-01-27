package priority

import (
	"testing"
)

func TestHeap(t *testing.T) {

	intInput := []uint8{4, 6, 2, 7, 9, 8, 3, 1, 5, 0}
	ascInts := []uint8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	desInts := []uint8{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}

	for _, test := range []*heapTest[uint8]{
		{"binary min heap uint8", MinHeapBinary(intInput...), ascInts},
		{"binary max heap uint8", MaxHeapBinary(intInput...), desInts},
		{"fib min heap uint8", MinHeapFib(intInput...), ascInts},
		{"fib max heap uint8", MinHeapFib(intInput...), desInts},
	} {
		test.run(t)
	}

	strInput := []string{"d", "a", "g", "c", "f", "b", "e"}

	for _, test := range []*heapTest[string]{
		{"min heap string", MinHeapBinary(strInput...), []string{"a", "b", "c", "d", "e", "f", "g"}},
		{"max heap string", MaxHeapBinary(strInput...), []string{"g", "f", "e", "d", "c", "b", "a"}},
	} {
		test.run(t)
	}
}

type heapTest[T comparable] struct {
	name     string
	pq       Queue[T]
	expected []T
}

func (ht *heapTest[T]) run(t *testing.T) {
	t.Run(ht.name, func(t *testing.T) {
		var actual []T
		for !ht.pq.IsEmpty() {
			top, _ := ht.pq.Pop()
			if peek, ok := ht.pq.Peek(); !ok || peek != top {
				t.Errorf("peek (%v) does not equal pop (%v)", peek, top)

			}
			actual = append(actual, top)
		}
		assertSlicesEqual(t, ht.expected, actual)
	})
}

func assertSlicesEqual[T comparable](t *testing.T, expected, actual []T) {
	if le, la := len(expected), len(actual); le != la {
		t.Errorf("exected slice of length %d - but got slice of length %d\n - %v\n - %v", le, la, expected, actual)
	}

	for i, v := range expected {
		if va := actual[i]; v != va {
			t.Errorf("mismatching values at index %d: expected %v but got %v", i, v, va)
		}
	}
}
