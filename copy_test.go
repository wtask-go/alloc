package alloc_test

import (
	"testing"
	"time"

	"github.com/wtask-go/alloc"
)

func makeTestCopy[T alloc.Scope](cases ...T) func(*testing.T) {
	return func(t *testing.T) {
		t.Helper()

		if ptr := alloc.Copy[T](nil); ptr != nil {
			t.Fatal("got pointer to:", *ptr, "expected: nil")
		}

		if len(cases) == 0 {
			t.Fatal("useless test: no cases")
		}

		for _, c := range cases {
			var (
				val = c
				ptr = &val
			)

			switch res := alloc.Copy(ptr); {
			case res == nil:
				t.Fatal("unexpected nil for:", val)
			case res == ptr:
				t.Fatal("got the same pointer")
			case *res != val:
				t.Fatal("referenced value mismatch:", "got:", *res, "expected:", val)
			}
		}
	}
}

//nolint:revive // function-length: it is require to test all type constraints provided by alloc.Scope
func TestCopy(t *testing.T) {
	t.Parallel()

	t.Run("untyped int", makeTestCopy(-1, 0, 1))
	t.Run("int", makeTestCopy[int](-1, 0, 1))
	t.Run("int8", makeTestCopy[int8](-1, 0, 1))
	t.Run("int16", makeTestCopy[int16](-1, 0, 1))
	t.Run("int32", makeTestCopy[int32](-1, 0, 1))
	t.Run("int64", makeTestCopy[int64](-1, 0, 1))
	t.Run("myInt", makeTestCopy(types.myInt...))
	t.Run("myInt8", makeTestCopy(types.myInt8...))
	t.Run("myInt16", makeTestCopy(types.myInt16...))
	t.Run("myInt32", makeTestCopy(types.myInt32...))
	t.Run("myInt64", makeTestCopy(types.myInt64...))

	t.Run("uint", makeTestCopy[uint](0, 1))
	t.Run("uint8", makeTestCopy[uint8](0, 1))
	t.Run("uint16", makeTestCopy[uint16](0, 1))
	t.Run("uint32", makeTestCopy[uint32](0, 1))
	t.Run("uint64", makeTestCopy[uint64](0, 1))
	t.Run("myUInt", makeTestCopy(types.myUInt...))
	t.Run("myUInt8", makeTestCopy(types.myUInt8...))
	t.Run("myUInt16", makeTestCopy(types.myUInt16...))
	t.Run("myUInt32", makeTestCopy(types.myUInt32...))
	t.Run("myUInt64", makeTestCopy(types.myUInt64...))

	t.Run("uintptr", makeTestCopy[uintptr](0, 1))
	t.Run("myUIntPtr", makeTestCopy(types.myUIntPtr...))

	t.Run("untyped float", makeTestCopy(-1.0, 0.0, 1.0))
	t.Run("float32", makeTestCopy[float32](-1.0, 0.0, 1.0))
	t.Run("float64", makeTestCopy[float64](-1.0, 0.0, 1.0))
	t.Run("myFloat32", makeTestCopy(types.myFloat32...))
	t.Run("myFloat64", makeTestCopy(types.myFloat64...))

	t.Run("untyped string", makeTestCopy("", "untyped string"))
	t.Run("string", makeTestCopy[string]("", "string"))
	t.Run("myString", makeTestCopy(types.myString...))

	t.Run("untyped bool", makeTestCopy(false, true))
	t.Run("bool", makeTestCopy[bool](false, true))
	t.Run("myBool", makeTestCopy[myBool](false, true))

	t.Run("time.Time", makeTestCopy(time.Time{}, time.Now()))

	t.Run("time.Duration", makeTestCopy(
		-1*time.Nanosecond,
		0*time.Nanosecond,
		1*time.Nanosecond,
		1*time.Microsecond,
		1*time.Millisecond,
		1*time.Second,
	))
}
