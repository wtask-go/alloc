package alloc_test

import (
	"testing"
	"time"

	"github.com/wtask-go/alloc"
)

func testNew[T alloc.Scope](cases ...T) func(*testing.T) {
	return func(t *testing.T) {
		t.Helper()

		if len(cases) == 0 {
			t.Fatal("useless test: no cases")
		}

		for _, c := range cases {
			switch ref := alloc.New(c); {
			case ref == nil:
				t.Fatal("nil allocation for:", c)
			case *ref != c:
				t.Fatal("got value:", *ref, "expected:", c)
			}
		}

	}
}

func TestNew(t *testing.T) {
	t.Parallel()

	t.Run("untyped int", testNew(-1, 0, 1))
	t.Run("int", testNew[int](-1, 0, 1))
	t.Run("int8", testNew[int8](-1, 0, 1))
	t.Run("int16", testNew[int16](-1, 0, 1))
	t.Run("int32", testNew[int32](-1, 0, 1))
	t.Run("int64", testNew[int64](-1, 0, 1))
	t.Run("myInt", testNew(types.myInt...))
	t.Run("myInt8", testNew(types.myInt8...))
	t.Run("myInt16", testNew(types.myInt16...))
	t.Run("myInt32", testNew(types.myInt32...))
	t.Run("myInt64", testNew(types.myInt64...))

	t.Run("uint", testNew[uint](0, 1))
	t.Run("uint8", testNew[uint8](0, 1))
	t.Run("uint16", testNew[uint16](0, 1))
	t.Run("uint32", testNew[uint32](0, 1))
	t.Run("uint64", testNew[uint64](0, 1))
	t.Run("myUInt", testNew(types.myUInt...))
	t.Run("myUInt8", testNew(types.myUInt8...))
	t.Run("myUInt16", testNew(types.myUInt16...))
	t.Run("myUInt32", testNew(types.myUInt32...))
	t.Run("myUInt64", testNew(types.myUInt64...))

	t.Run("uintptr", testNew[uintptr](0, 1))
	t.Run("myUIntPtr", testNew(types.myUIntPtr...))

	t.Run("untyped float", testNew(-1.0, 0.0, 1.0))
	t.Run("float32", testNew[float32](-1.0, 0.0, 1.0))
	t.Run("float64", testNew[float64](-1.0, 0.0, 1.0))
	t.Run("myFloat32", testNew(types.myFloat32...))
	t.Run("myFloat64", testNew(types.myFloat64...))

	t.Run("untyped string", testNew("", "untyped string"))
	t.Run("string", testNew[string]("", "string"))
	t.Run("myString", testNew(types.myString...))

	t.Run("time.Time", testNew(time.Time{}, time.Now()))
}
