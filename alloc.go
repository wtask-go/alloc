// Package alloc is intended to help make pointers to specified values with single call
// and also to get values from pointers safely to avoid nil dereference errors.
package alloc

import (
	"time"

	"golang.org/x/exp/constraints"
)

// Scope represents type constraints to allow generalized operations within this package.
type Scope interface {
	constraints.Ordered | ~bool | time.Time | time.Duration
}

// New is generic function to allocate new pointer to specified value.
func New[V Scope](value V) *V {
	return &value
}

// Deref is generic function to dereference a pointer.
// Returns referenced value and true for non nil pointer or type zero value and false otherwise.
func Deref[V Scope](ref *V) (V, bool) {
	var zero V

	if ref == nil {
		return zero, false
	}

	return *ref, true
}

// Value is generic function to dereference a pointer.
// Returns referenced value for non nil pointer or type zero value otherwise.
func Value[V Scope](ref *V) V {
	v, _ := Deref(ref)

	return v
}

// Copy is generic function to duplicate specified pointer with referenced value.
// Returns new pointer for specified non nil reference or nil otherwise.
func Copy[V Scope](ref *V) *V {
	if ref == nil {
		return nil
	}

	return New(Value(ref))
}
