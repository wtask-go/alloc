package alloc_test

import (
	"testing"
	"time"

	"github.com/wtask-go/alloc"
)

func testDeref[T alloc.Scope](cases ...T) func(*testing.T) {
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

func TestDeref(t *testing.T) {
	t.Parallel()

	t.Run("untyped int", testDeref(-1, 0, 1))
	t.Run("int", testDeref[int](-1, 0, 1))
	t.Run("int8", testDeref[int8](-1, 0, 1))
	t.Run("int16", testDeref[int16](-1, 0, 1))
	t.Run("int32", testDeref[int32](-1, 0, 1))
	t.Run("int64", testDeref[int64](-1, 0, 1))
	t.Run("myInt", testDeref(types.myInt...))
	t.Run("myInt8", testDeref(types.myInt8...))
	t.Run("myInt16", testDeref(types.myInt16...))
	t.Run("myInt32", testDeref(types.myInt32...))
	t.Run("myInt64", testDeref(types.myInt64...))

	t.Run("uint", testDeref[uint](0, 1))
	t.Run("uint8", testDeref[uint8](0, 1))
	t.Run("uint16", testDeref[uint16](0, 1))
	t.Run("uint32", testDeref[uint32](0, 1))
	t.Run("uint64", testDeref[uint64](0, 1))
	t.Run("myUInt", testDeref(types.myUInt...))
	t.Run("myUInt8", testDeref(types.myUInt8...))
	t.Run("myUInt16", testDeref(types.myUInt16...))
	t.Run("myUInt32", testDeref(types.myUInt32...))
	t.Run("myUInt64", testDeref(types.myUInt64...))

	t.Run("uintptr", testDeref[uintptr](0, 1))
	t.Run("myUIntPtr", testDeref(types.myUIntPtr...))

	t.Run("untyped float", testDeref(-1.0, 0.0, 1.0))
	t.Run("float32", testDeref[float32](-1.0, 0.0, 1.0))
	t.Run("float64", testDeref[float64](-1.0, 0.0, 1.0))
	t.Run("myFloat32", testDeref(types.myFloat32...))
	t.Run("myFloat64", testDeref(types.myFloat64...))

	t.Run("untyped string", testDeref("", "untyped string"))
	t.Run("string", testDeref[string]("", "string"))
	t.Run("myString", testDeref(types.myString...))

	t.Run("untyped bool", testDeref(false, true))
	t.Run("bool", testDeref[bool](false, true))
	t.Run("myBool", testDeref[myBool](false, true))

	t.Run("time.Time", testDeref(time.Time{}, time.Now()))
}
