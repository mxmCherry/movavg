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
