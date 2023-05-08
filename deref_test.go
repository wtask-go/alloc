package alloc_test

import (
	"testing"
	"time"

	"github.com/wtask-go/alloc"
)

func makeTestDeref[T alloc.Scope](cases ...T) func(*testing.T) {
	return func(t *testing.T) {
		t.Helper()

		var zero T

		// nil dereference
		switch val, ok := alloc.Deref[T](nil); {
		case ok:
			t.Fatal("nil dereference:", "ok=true")
		case val != zero:
			t.Fatal("nil dereference:", "got non-zero value:", val)
		}

		if len(cases) == 0 {
			t.Error("incomplete test:", "no cases")
		}

		for _, c := range cases {
			switch val, ok := alloc.Deref(&c); {
			case !ok:
				t.Fatal("value dereference:", "ok=false")
			case val != c:
				t.Fatal("value dereference:", "got:", val, "expected:", c)
			}
		}
	}
}

//nolint:revive // function-length: it is required to test all type constraints provided by alloc.Scope
func TestDeref(t *testing.T) {
	t.Parallel()

	t.Run("untyped int", makeTestDeref(-1, 0, 1))
	t.Run("int", makeTestDeref[int](-1, 0, 1))
	t.Run("int8", makeTestDeref[int8](-1, 0, 1))
	t.Run("int16", makeTestDeref[int16](-1, 0, 1))
	t.Run("int32", makeTestDeref[int32](-1, 0, 1))
	t.Run("int64", makeTestDeref[int64](-1, 0, 1))
	t.Run("myInt", makeTestDeref(types.myInt...))
	t.Run("myInt8", makeTestDeref(types.myInt8...))
	t.Run("myInt16", makeTestDeref(types.myInt16...))
	t.Run("myInt32", makeTestDeref(types.myInt32...))
	t.Run("myInt64", makeTestDeref(types.myInt64...))

	t.Run("uint", makeTestDeref[uint](0, 1))
	t.Run("uint8", makeTestDeref[uint8](0, 1))
	t.Run("uint16", makeTestDeref[uint16](0, 1))
	t.Run("uint32", makeTestDeref[uint32](0, 1))
	t.Run("uint64", makeTestDeref[uint64](0, 1))
	t.Run("myUInt", makeTestDeref(types.myUInt...))
	t.Run("myUInt8", makeTestDeref(types.myUInt8...))
	t.Run("myUInt16", makeTestDeref(types.myUInt16...))
	t.Run("myUInt32", makeTestDeref(types.myUInt32...))
	t.Run("myUInt64", makeTestDeref(types.myUInt64...))

	t.Run("uintptr", makeTestDeref[uintptr](0, 1))
	t.Run("myUIntPtr", makeTestDeref(types.myUIntPtr...))

	t.Run("untyped float", makeTestDeref(-1.0, 0.0, 1.0))
	t.Run("float32", makeTestDeref[float32](-1.0, 0.0, 1.0))
	t.Run("float64", makeTestDeref[float64](-1.0, 0.0, 1.0))
	t.Run("myFloat32", makeTestDeref(types.myFloat32...))
	t.Run("myFloat64", makeTestDeref(types.myFloat64...))

	t.Run("untyped string", makeTestDeref("", "untyped string"))
	t.Run("string", makeTestDeref[string]("", "string"))
	t.Run("myString", makeTestDeref(types.myString...))

	t.Run("untyped bool", makeTestDeref(false, true))
	t.Run("bool", makeTestDeref[bool](false, true))
	t.Run("myBool", makeTestDeref[myBool](false, true))

	t.Run("time.Time", makeTestDeref(time.Time{}, time.Now()))

	t.Run("time.Duration", makeTestDeref(
		-1*time.Nanosecond,
		0*time.Nanosecond,
		1*time.Nanosecond,
		1*time.Microsecond,
		1*time.Millisecond,
		1*time.Second,
	))
}
