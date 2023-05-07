package alloc_test

import (
	"testing"
	"time"

	"github.com/wtask-go/alloc"
)

func testCopy[T alloc.Scope](cases ...T) func(*testing.T) {
	return func(t *testing.T) {
		t.Helper()

		if ptr := alloc.Copy[T](nil); ptr != nil {
			t.Fatal("got pointer to:", *ptr, "expected: nil")
		}

		if len(cases) == 0 {
			t.Fatal("useless test: no cases")
		}

		for _, c := range cases {
			ptr := &c
			switch res := alloc.Copy(ptr); {
			case res == nil:
				t.Fatal("unexpected nil for:", c)
			case res == ptr:
				t.Fatal("got the same pointer")
			case *res != c:
				t.Fatal("referenced value mismatch:", "got:", *res, "expected:", c)
			}
		}
	}
}

func TestCopy(t *testing.T) {
	t.Parallel()

	t.Run("untyped int", testCopy(-1, 0, 1))
	t.Run("int", testCopy[int](-1, 0, 1))
	t.Run("int8", testCopy[int8](-1, 0, 1))
	t.Run("int16", testCopy[int16](-1, 0, 1))
	t.Run("int32", testCopy[int32](-1, 0, 1))
	t.Run("int64", testCopy[int64](-1, 0, 1))
	t.Run("myInt", testCopy(types.myInt...))
	t.Run("myInt8", testCopy(types.myInt8...))
	t.Run("myInt16", testCopy(types.myInt16...))
	t.Run("myInt32", testCopy(types.myInt32...))
	t.Run("myInt64", testCopy(types.myInt64...))

	t.Run("uint", testCopy[uint](0, 1))
	t.Run("uint8", testCopy[uint8](0, 1))
	t.Run("uint16", testCopy[uint16](0, 1))
	t.Run("uint32", testCopy[uint32](0, 1))
	t.Run("uint64", testCopy[uint64](0, 1))
	t.Run("myUInt", testCopy(types.myUInt...))
	t.Run("myUInt8", testCopy(types.myUInt8...))
	t.Run("myUInt16", testCopy(types.myUInt16...))
	t.Run("myUInt32", testCopy(types.myUInt32...))
	t.Run("myUInt64", testCopy(types.myUInt64...))

	t.Run("uintptr", testCopy[uintptr](0, 1))
	t.Run("myUIntPtr", testCopy(types.myUIntPtr...))

	t.Run("untyped float", testCopy(-1.0, 0.0, 1.0))
	t.Run("float32", testCopy[float32](-1.0, 0.0, 1.0))
	t.Run("float64", testCopy[float64](-1.0, 0.0, 1.0))
	t.Run("myFloat32", testCopy(types.myFloat32...))
	t.Run("myFloat64", testCopy(types.myFloat64...))

	t.Run("untyped string", testCopy("", "untyped string"))
	t.Run("string", testCopy[string]("", "string"))
	t.Run("myString", testCopy(types.myString...))

	t.Run("untyped bool", testCopy(false, true))
	t.Run("bool", testCopy[bool](false, true))
	t.Run("myBool", testCopy[myBool](false, true))

	t.Run("time.Time", testCopy(time.Time{}, time.Now()))
}
