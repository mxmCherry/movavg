package movavg_test

import (
	"fmt"

	. "github.com/mxmCherry/movavg"
)

func ExampleMulti() {

	multiMA := MultiThreadSafe(
		Multi{
			NewSMA(2),
			NewSMA(3),
			NewSMA(4),
		},
	)

	// zero until values added:
	fmt.Println(multiMA.Avg()) // [0 0 0]

	// load initial value
	fmt.Println(multiMA.Add(2)) // [2 2 2]
	fmt.Println(multiMA.Avg())  // [2 2 2]

	fmt.Println(multiMA.Add(4)) // [3 3 3]
	fmt.Println(multiMA.Avg())  // [3 3 3]

	fmt.Println(multiMA.Add(6)) // [5 4 4]
	fmt.Println(multiMA.Avg())  // [5 4 4]

	fmt.Println(multiMA.Add(8)) // [7 6 5]
	fmt.Println(multiMA.Avg())  // [7 6 5]

	fmt.Println(multiMA.Add(10)) // [9 8 7]
	fmt.Println(multiMA.Avg())   // [9 8 7]

	// Output:
	// [0 0 0]
	// [2 2 2]
	// [2 2 2]
	// [3 3 3]
	// [3 3 3]
	// [5 4 4]
	// [5 4 4]
	// [7 6 5]
	// [7 6 5]
	// [9 8 7]
	// [9 8 7]
}
