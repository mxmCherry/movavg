# movavg [![GoDoc](https://godoc.org/github.com/mxmCherry/movavg?status.svg)](https://godoc.org/github.com/mxmCherry/movavg) [![Build Status](https://travis-ci.org/mxmCherry/movavg.svg?branch=master)](https://travis-ci.org/mxmCherry/movavg) [![Go Report Card](https://goreportcard.com/badge/github.com/mxmCherry/movavg)](https://goreportcard.com/report/github.com/mxmCherry/movavg)

[Moving Average](https://en.wikipedia.org/wiki/Moving_average) calculator for Go (Golang)

## Example

```go
package movavg_test

import (
	"fmt"

	. "github.com/mxmCherry/movavg"
)

func ExampleSMA() {

	sma := ThreadSafe(NewSMA(3))

	fmt.Println(sma.Avg())

	fmt.Println(sma.Add(1))
	fmt.Println(sma.Avg())

	fmt.Println(sma.Add(2))
	fmt.Println(sma.Avg())

	fmt.Println(sma.Add(3))
	fmt.Println(sma.Avg())

	fmt.Println(sma.Add(4))
	fmt.Println(sma.Avg())

	fmt.Println(sma.Add(5))
	fmt.Println(sma.Avg())

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
