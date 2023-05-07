// Package alloc is intended to help make pointers to specified values with single call
// and also to get values from pointers safely to avoid nil dereference errors.
package alloc

import (
	"time"

	"golang.org/x/exp/constraints"
)

type Scope interface {
	constraints.Ordered | ~bool | time.Time | time.Duration
}

func New[V Scope](value V) *V {
	return &value
}

func Deref[V Scope](ref *V) (V, bool) {
	var zero V

	if ref == nil {
		return zero, false
	}

	return *ref, true
}

func Value[V Scope](ref *V) V {
	v, _ := Deref(ref)

	return v
}

func Copy[V Scope](ref *V) *V {
	if ref == nil {
		return nil
	}

	return New(Value(ref))
}
