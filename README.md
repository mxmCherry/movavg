# movavg [![GoDoc](https://godoc.org/github.com/mxmCherry/movavg?status.svg)](https://godoc.org/github.com/mxmCherry/movavg) [![Build Status](https://travis-ci.org/mxmCherry/movavg.svg?branch=master)](https://travis-ci.org/mxmCherry/movavg) [![Go Report Card](https://goreportcard.com/badge/github.com/mxmCherry/movavg)](https://goreportcard.com/report/github.com/mxmCherry/movavg)

[Moving Average](https://en.wikipedia.org/wiki/Moving_average) calculator for Go (Golang) with focus on performance.

Pros:
- math instead of bruteforce loops for recalc (roughly, `O(1)` recalcs)
- zero allocations (only in constructors, not on recalcs)
- thread-safety "on demand" (`ThreadSafe(MA)` adapter)

Cons:
- `float64` is not very precise, floating-point errors may be accumulated

## Example

```go
package movavg_test

import (
	"fmt"

	. "github.com/mxmCherry/movavg"
)

func ExampleSMA() {

	sma := ThreadSafe(NewSMA(3))

	// zero until values added:
	fmt.Println(sma.Avg()) // 0

	fmt.Println(sma.Add(1)) // 1
	fmt.Println(sma.Avg())  // 1

	fmt.Println(sma.Add(2)) // 1.5
	fmt.Println(sma.Avg())  // 1.5

	fmt.Println(sma.Add(3)) // 2
	fmt.Println(sma.Avg())  // 2

	fmt.Println(sma.Add(4)) // 3
	fmt.Println(sma.Avg())  // 3

	fmt.Println(sma.Add(5)) // 4
	fmt.Println(sma.Avg())  // 4

	// Output:
	// 0
	// 1
	// 1
	// 1.5
	// 1.5
	// 2
	// 2
	// 3
	// 3
	// 4
	// 4
}
```
