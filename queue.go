// Package priority provides generic priority queue implementations
package priority

// PQ provides a priority queue interface
type Queue[T any] interface {
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

// predicate provides a function to determine the priority between two items
// returns true if a has a higher priority that b, false otherwise
type predicate[T any] func(a, b T) bool
