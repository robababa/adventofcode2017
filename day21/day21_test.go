package main

import (
"fmt"
)

// 1 2  => 3 4
// 3 4  => 1 2
// 1 2 3 => 7 8 9
// 4 5 6 => 4 5 6
// 7 8 9 => 1 2 3
func ExampleFlipTopAndBottom() {
	fmt.Println(flipTopAndBottom([]string{"12","34"}))
	fmt.Println(flipTopAndBottom([]string{"123","456","789"}))
	//Output:
	//[34 12]
	//[789 456 123]
}

// 1 2  => 2 1
// 3 4  => 4 3
// 1 2 3 => 3 2 1
// 4 5 6 => 6 5 4
// 7 8 9 => 9 8 7
func ExampleFlipLeftAndRight() {
	fmt.Println(flipLeftAndRight([]string{"12","34"}))
	fmt.Println(flipLeftAndRight([]string{"123","456","789"}))
	//Output:
	//[21 43]
	//[321 654 987]
}

// 1 2  => 3 1
// 3 4  => 4 2
// 1 2 3 => 7 4 1
// 4 5 6 => 8 5 2
// 7 8 9 => 9 6 3
func ExampleRotateClockwise() {
	fmt.Println(rotateClockwise([]string{"12","34"}))
	fmt.Println(rotateClockwise([]string{"123","456","789"}))
	//Output:
	//[31 42]
	//[741 852 963]
}