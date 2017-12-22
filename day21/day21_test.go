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
func ExampleRotate() {
	fmt.Println(rotate([]string{"12","34"}, 1))
	fmt.Println(rotate([]string{"12","34"}, 2))
	fmt.Println(rotate([]string{"12","34"}, 3))
	fmt.Println(rotate([]string{"123","456","789"}, 1))
	fmt.Println(rotate([]string{"123","456","789"}, 2))
	fmt.Println(rotate([]string{"123","456","789"}, 3))
	//Output:
	//[31 42]
	//[43 21]
	//[24 13]
	//[741 852 963]
	//[987 654 321]
	//[369 258 147]
}

func ExampleGridToKey() {
	fmt.Println(gridToKey([]string{"12", "34"}))
	fmt.Println(gridToKey([]string{"123", "456", "789"}))
	//Output:
	//12/34
	//123/456/789
}

//  1  2 | 5  6
//  3  4 | 7  8
// --------------
//  A  B | E  F
//  C  D | G  H
func ExampleCombineGrids() {
	fmt.Println(combineGrids([][]string{{"12","34"}}))
	fmt.Println(combineGrids([][]string{{"12","34"},{"56","78"},{"AB","CD"},{"EF","GH"}}))
	//Output:
	//[12 34]
	//[1256 3478 ABEF CDGH]
}

// 12
// 34
//
// 123
// 456
// 789
//
// 1234
// 5678
// ABCD
// EFGH
func ExampleDivideGrid() {
	fmt.Println(divideGrid([]string{"12","34"}))
	fmt.Println(divideGrid([]string{"123","456","789"}))
	fmt.Println(divideGrid([]string{"1234","5678","ABCD","EFGH"}))
	//Output:
	//[[12 34]]
	//[[123 456 789]]
	//[[12 56] [34 78] [AB EF] [CD GH]]
}