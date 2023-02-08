// Package priority provides generic priority queue implementations
package priority

import "golang.org/x/exp/constraints"

// Queue provides a priority queue interface
type Queue[T comparable] interface {

	// IsEmpty returns true if the priority queue is empty - false otherwise
	IsEmpty() bool

	// Peek provides the highest priority item without removing it from the priority queue
	// boolean retured indicates if the operation was successful i.e returns false if structure is empty
	Peek() (T, bool)

	// Pop provides the highest priority item and removes it from the priority queue
	// boolean retured indicates if the operation was successful i.e returns false if structure is empty
	Pop() (T, bool)

	// Push adds an item to the priority queue
	Push(T)
}

type QueueWithUpdate[T comparable] interface {
	Queue[T]

	Update(T, UpdateFunc[T])
}

// predicate provides a function to determine the priority between two items
// returns true if a has a higher priority that b, false otherwise
type predicate[T any] func(a, b T) bool

func lessThan[T constraints.Ordered](a, b T) bool {
	return a < b
}

func greaterThan[T constraints.Ordered](a, b T) bool {
	return a > b
}

type UpdateFunc[T any] func(T) T
