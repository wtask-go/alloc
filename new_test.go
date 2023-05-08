package alloc_test

import (
	"testing"
	"time"

	"github.com/wtask-go/alloc"
)

func makeTestNew[T alloc.Scope](cases ...T) func(*testing.T) {
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

//nolint:revive // function-length: it is required to test all type constraints provided by alloc.Scope
func TestNew(t *testing.T) {
	t.Parallel()

	t.Run("untyped int", makeTestNew(-1, 0, 1))
	t.Run("int", makeTestNew[int](-1, 0, 1))
	t.Run("int8", makeTestNew[int8](-1, 0, 1))
	t.Run("int16", makeTestNew[int16](-1, 0, 1))
	t.Run("int32", makeTestNew[int32](-1, 0, 1))
	t.Run("int64", makeTestNew[int64](-1, 0, 1))
	t.Run("myInt", makeTestNew(types.myInt...))
	t.Run("myInt8", makeTestNew(types.myInt8...))
	t.Run("myInt16", makeTestNew(types.myInt16...))
	t.Run("myInt32", makeTestNew(types.myInt32...))
	t.Run("myInt64", makeTestNew(types.myInt64...))

	t.Run("uint", makeTestNew[uint](0, 1))
	t.Run("uint8", makeTestNew[uint8](0, 1))
	t.Run("uint16", makeTestNew[uint16](0, 1))
	t.Run("uint32", makeTestNew[uint32](0, 1))
	t.Run("uint64", makeTestNew[uint64](0, 1))
	t.Run("myUInt", makeTestNew(types.myUInt...))
	t.Run("myUInt8", makeTestNew(types.myUInt8...))
	t.Run("myUInt16", makeTestNew(types.myUInt16...))
	t.Run("myUInt32", makeTestNew(types.myUInt32...))
	t.Run("myUInt64", makeTestNew(types.myUInt64...))

	t.Run("uintptr", makeTestNew[uintptr](0, 1))
	t.Run("myUIntPtr", makeTestNew(types.myUIntPtr...))

	t.Run("untyped float", makeTestNew(-1.0, 0.0, 1.0))
	t.Run("float32", makeTestNew[float32](-1.0, 0.0, 1.0))
	t.Run("float64", makeTestNew[float64](-1.0, 0.0, 1.0))
	t.Run("myFloat32", makeTestNew(types.myFloat32...))
	t.Run("myFloat64", makeTestNew(types.myFloat64...))

	t.Run("untyped string", makeTestNew("", "untyped string"))
	t.Run("string", makeTestNew[string]("", "string"))
	t.Run("myString", makeTestNew(types.myString...))

	t.Run("untyped bool", makeTestNew(false, true))
	t.Run("bool", makeTestNew[bool](false, true))
	t.Run("myBool", makeTestNew[myBool](false, true))

	t.Run("time.Time", makeTestNew(time.Time{}, time.Now()))

	t.Run("time.Duration", makeTestNew(
		-1*time.Nanosecond,
		0*time.Nanosecond,
		1*time.Nanosecond,
		1*time.Microsecond,
		1*time.Millisecond,
		1*time.Second,
	))
}
