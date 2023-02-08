package priority

import (
	"fmt"
	"testing"
)

var (
	intInput = []uint8{4, 6, 2, 7, 9, 8, 3, 1, 5, 0}
	ascInts  = []uint8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	desInts  = []uint8{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	strInput = []string{"d", "a", "g", "c", "f", "b", "e"}
)

func TestHeap(t *testing.T) {

	for _, test := range []*heapTest[uint8]{
		{"linear min pq uint8", MinLinearQueue(intInput...), ascInts},
		{"linear max pq uint8", MaxLinearQueue(intInput...), desInts},
		{"binary min heap uint8", MinHeapBinary(intInput...), ascInts},
		{"binary max heap uint8", MaxHeapBinary(intInput...), desInts},
		{"binomial min heap uint8", MinHeapBinary(intInput...), ascInts},
		{"binomial max heap uint8", MaxHeapBinary(intInput...), desInts},
		// {"fib min heap uint8", MinHeapFib(intInput...), ascInts},
		// {"fib max heap uint8", MinHeapFib(intInput...), desInts},
	} {
		test.run(t)
	}

	for _, test := range []*heapTest[string]{
		{"min heap string", MinHeapBinary(strInput...), []string{"a", "b", "c", "d", "e", "f", "g"}},
		{"max heap string", MaxHeapBinary(strInput...), []string{"g", "f", "e", "d", "c", "b", "a"}},
	} {
		test.run(t)
	}
}

func TestDecreaseKey(t *testing.T) {
	for _, test := range []struct {
		name     string
		pq       QueueWithUpdate[uint8]
		update   UpdateFunc[uint8]
		expected []uint8
	}{
		{"binary min heap uint8 update to square", MinHeapBinary(intInput...), squareUint8, updatedArray(ascInts, squareUint8)},
		{"binary max heap uint8 update to square", MaxHeapBinary(intInput...), squareUint8, updatedArray(desInts, squareUint8)},
		{"binary min heap uint8 update to double", MinHeapBinary(intInput...), doubleUint8, updatedArray(ascInts, doubleUint8)},
		{"binary max heap uint8 update to double", MaxHeapBinary(intInput...), doubleUint8, updatedArray(desInts, doubleUint8)},
	} {
		test := test // pin
		t.Run(test.name, func(t *testing.T) {
			for _, i := range intInput {
				test.pq.Update(i, test.update)
			}

			var actual []uint8
			for !test.pq.IsEmpty() {
				top, _ := test.pq.Pop()
				actual = append(actual, top)
			}

			assertSlicesEqual(t, test.expected, actual)
		})
	}
}

func squareUint8(i uint8) uint8 {
	return i * i
}

func doubleUint8(i uint8) uint8 {
	return i * 2
}

func updatedArray(arr []uint8, f UpdateFunc[uint8]) []uint8 {
	result := make([]uint8, len(arr))

	for i, n := range arr {
		result[i] = f(n)
	}

	return result
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
			top, _ := ht.pq.Peek()
			if pop, ok := ht.pq.Pop(); !ok || pop != top {
				t.Errorf("peek (%v) does not equal pop (%v)", top, pop)
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

	fmt.Println(actual)
	for i, v := range expected {
		if va := actual[i]; v != va {
			t.Errorf("mismatching values at index %d: expected %v but got %v", i, v, va)
		}
	}
}
