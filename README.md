# tysyncmap

[![Go Reference]([https://pkg.go.dev/badge/github.com/umgefahren/tysyncmap.svg](https://pkg.go.dev/badge/github.com/umgefahren/tysyncmap.svg))]([https://pkg.go.dev/github.com/umgefahren/tysyncmap](https://pkg.go.dev/github.com/umgefahren/tysyncmap))
Typed version of sync.Map using Go 1.18 generics.

## Examples

```golang
package main

import (
	"fmt"
	"tysyncmap"
)

func main() {
	m := new(tysyncmap.Map[string, string])
	m.Store("key", "value")
	val, ok := m.Load("key")
	if !ok {
		panic("should have been loaded")
	}
	fmt.Println(val)
}

```

 Output:

```
value
```
