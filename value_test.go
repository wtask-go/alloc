package alloc_test

import (
	"testing"
	"time"

	"github.com/wtask-go/alloc"
)

func testValue[T alloc.Scope](cases ...T) func(*testing.T) {
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

func TestValue(t *testing.T) {
	t.Parallel()

	t.Run("untyped int", testValue(-1, 0, 1))
	t.Run("int", testValue[int](-1, 0, 1))
	t.Run("int8", testValue[int8](-1, 0, 1))
	t.Run("int16", testValue[int16](-1, 0, 1))
	t.Run("int32", testValue[int32](-1, 0, 1))
	t.Run("int64", testValue[int64](-1, 0, 1))
	t.Run("myInt", testValue(types.myInt...))
	t.Run("myInt8", testValue(types.myInt8...))
	t.Run("myInt16", testValue(types.myInt16...))
	t.Run("myInt32", testValue(types.myInt32...))
	t.Run("myInt64", testValue(types.myInt64...))

	t.Run("uint", testValue[uint](0, 1))
	t.Run("uint8", testValue[uint8](0, 1))
	t.Run("uint16", testValue[uint16](0, 1))
	t.Run("uint32", testValue[uint32](0, 1))
	t.Run("uint64", testValue[uint64](0, 1))
	t.Run("myUInt", testValue(types.myUInt...))
	t.Run("myUInt8", testValue(types.myUInt8...))
	t.Run("myUInt16", testValue(types.myUInt16...))
	t.Run("myUInt32", testValue(types.myUInt32...))
	t.Run("myUInt64", testValue(types.myUInt64...))

	t.Run("uintptr", testValue[uintptr](0, 1))
	t.Run("myUIntPtr", testValue(types.myUIntPtr...))

	t.Run("untyped float", testValue(-1.0, 0.0, 1.0))
	t.Run("float32", testValue[float32](-1.0, 0.0, 1.0))
	t.Run("float64", testValue[float64](-1.0, 0.0, 1.0))
	t.Run("myFloat32", testValue(types.myFloat32...))
	t.Run("myFloat64", testValue(types.myFloat64...))

	t.Run("untyped string", testValue("", "untyped string"))
	t.Run("string", testValue[string]("", "string"))
	t.Run("myString", testValue(types.myString...))

	t.Run("time.Time", testValue(time.Time{}, time.Now()))
}
