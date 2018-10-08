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
