package alloc_test

type (
	myInt   int
	myInt8  int8
	myInt16 int16
	myInt32 int32
	myInt64 int64

	myUInt   uint
	myUInt8  uint8
	myUInt16 uint16
	myUInt32 uint32
	myUInt64 uint64

	myUIntPtr uintptr

	myFloat32 float32
	myFloat64 float64

	myString string

	myBool bool
)

// types contains extra data to test alloc-package
var types = struct {
	myInt   []myInt
	myInt8  []myInt8
	myInt16 []myInt16
	myInt32 []myInt32
	myInt64 []myInt64

	myUInt   []myUInt
	myUInt8  []myUInt8
	myUInt16 []myUInt16
	myUInt32 []myUInt32
	myUInt64 []myUInt64

	myUIntPtr []myUIntPtr

	myFloat32 []myFloat32
	myFloat64 []myFloat64

	myString []myString

	myBool []myBool
}{
	myInt:   []myInt{-1, 0, 1},
	myInt8:  []myInt8{-1, 0, 1},
	myInt16: []myInt16{-1, 0, 1},
	myInt32: []myInt32{-1, 0, 1},
	myInt64: []myInt64{-1, 0, 1},

	myUInt:   []myUInt{0, 1},
	myUInt8:  []myUInt8{0, 1},
	myUInt16: []myUInt16{0, 1},
	myUInt32: []myUInt32{0, 1},
	myUInt64: []myUInt64{0, 1},

	myUIntPtr: []myUIntPtr{0, 1},

	myFloat32: []myFloat32{-1.1, 0, 1.1},
	myFloat64: []myFloat64{-1.1, 0, 1.1},

	myString: []myString{"", "my string"},

	myBool: []myBool{false, true},
}
