package alloc_test

import (
	"testing"
	"time"

	"github.com/wtask-go/alloc"
)

func makeTestValue[T alloc.Scope](cases ...T) func(*testing.T) {
	return func(t *testing.T) {
		t.Helper()

		var zero T

		// nil dereference
		if val := alloc.Value[T](nil); val != zero {
			t.Fatal("nil dereference:", "got non-zero value:", val)
		}

		if len(cases) == 0 {
			t.Error("incomplete test:", "no cases")
		}

		for _, c := range cases {
			if val := alloc.Value(&c); val != c {
				t.Fatal("value dereference:", "got:", val, "expected:", c)
			}
		}
	}
}

//nolint:revive // function-length: it is required to test all type constraints provided by alloc.Scope
func TestValue(t *testing.T) {
	t.Parallel()

	t.Run("untyped int", makeTestValue(-1, 0, 1))
	t.Run("int", makeTestValue[int](-1, 0, 1))
	t.Run("int8", makeTestValue[int8](-1, 0, 1))
	t.Run("int16", makeTestValue[int16](-1, 0, 1))
	t.Run("int32", makeTestValue[int32](-1, 0, 1))
	t.Run("int64", makeTestValue[int64](-1, 0, 1))
	t.Run("myInt", makeTestValue(types.myInt...))
	t.Run("myInt8", makeTestValue(types.myInt8...))
	t.Run("myInt16", makeTestValue(types.myInt16...))
	t.Run("myInt32", makeTestValue(types.myInt32...))
	t.Run("myInt64", makeTestValue(types.myInt64...))

	t.Run("uint", makeTestValue[uint](0, 1))
	t.Run("uint8", makeTestValue[uint8](0, 1))
	t.Run("uint16", makeTestValue[uint16](0, 1))
	t.Run("uint32", makeTestValue[uint32](0, 1))
	t.Run("uint64", makeTestValue[uint64](0, 1))
	t.Run("myUInt", makeTestValue(types.myUInt...))
	t.Run("myUInt8", makeTestValue(types.myUInt8...))
	t.Run("myUInt16", makeTestValue(types.myUInt16...))
	t.Run("myUInt32", makeTestValue(types.myUInt32...))
	t.Run("myUInt64", makeTestValue(types.myUInt64...))

	t.Run("uintptr", makeTestValue[uintptr](0, 1))
	t.Run("myUIntPtr", makeTestValue(types.myUIntPtr...))

	t.Run("untyped float", makeTestValue(-1.0, 0.0, 1.0))
	t.Run("float32", makeTestValue[float32](-1.0, 0.0, 1.0))
	t.Run("float64", makeTestValue[float64](-1.0, 0.0, 1.0))
	t.Run("myFloat32", makeTestValue(types.myFloat32...))
	t.Run("myFloat64", makeTestValue(types.myFloat64...))

	t.Run("untyped string", makeTestValue("", "untyped string"))
	t.Run("string", makeTestValue[string]("", "string"))
	t.Run("myString", makeTestValue(types.myString...))

	t.Run("untyped bool", makeTestValue(false, true))
	t.Run("bool", makeTestValue[bool](false, true))
	t.Run("myBool", makeTestValue[myBool](false, true))

	t.Run("time.Time", makeTestValue(time.Time{}, time.Now()))

	t.Run("time.Duration", makeTestValue(
		-1*time.Nanosecond,
		0*time.Nanosecond,
		1*time.Nanosecond,
		1*time.Microsecond,
		1*time.Millisecond,
		1*time.Second,
	))
}
