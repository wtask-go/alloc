package alloc_test

import (
	"fmt"
	"time"

	"github.com/wtask-go/alloc"
)

// fictional data model
type model struct {
	Value    *float64
	Modified *time.Time
}

func (m *model) Print() {
	if m == nil {
		fmt.Println("model: <nil>")
		return
	}

	fmt.Printf("model: {%v, %v}\n", alloc.Value(m.Value), alloc.Value(m.Modified))
}

const defaultValue = 3.14

func Example_use() {
	m := &model{
		Value: alloc.New(defaultValue), // useful referencing to constant values
	}

	m.Print()

	m.Value = alloc.New(alloc.Value(m.Value) * 2) // inline operations with referenced values
	m.Modified = alloc.New(time.Date(2023, 5, 8, 15, 0, 0, 0, time.UTC))

	m.Print()

	val := alloc.Copy(m.Value)
	*val = defaultValue // m.value and val are referenced to different values
	fmt.Printf("m.value: %v, val: %v\n", alloc.Value(m.Value), alloc.Value(val))

	m.Value = nil
	if _, ok := alloc.Deref(m.Value); !ok {
		// synthetic case due to you able to directly check m.Value == nil
		fmt.Println("update is not required: nil value")
	}

	// pipelined computation with pointers

	if v := alloc.Value[float64](nil); v != 0 {
		fmt.Println("division result:", defaultValue/v)
	} else {
		fmt.Println("invalid division by zero")
	}

	fmt.Println("subtraction:", defaultValue-alloc.Value[float64](nil))

	// Output:
	// model: {3.14, 0001-01-01 00:00:00 +0000 UTC}
	// model: {6.28, 2023-05-08 15:00:00 +0000 UTC}
	// m.value: 6.28, val: 3.14
	// update is not required: nil value
	// invalid division by zero
	// subtraction: 3.14
}
