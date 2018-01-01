package main

import "fmt"

func ExampleTransformLine() {
	fmt.Println(transformLine("/dev/grid/node-x1-y10    90T   71T    19T   78%"))
	// Output:
	// 1 10 90 71 19 78%
}
