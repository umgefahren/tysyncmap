package tysyncmap_test

import (
	"fmt"
	"tysyncmap"
)

func Example() {
	m := new(tysyncmap.Map[string, string])
	m.Store("key", "value")
	val, ok := m.Load("key")
	if !ok {
		panic("should have been loaded")
	}
	fmt.Println(val)
	// Output: value
}
