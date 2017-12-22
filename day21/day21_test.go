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
	fmt.Println(rotate([]string{"12","34"}, 1, true))
	fmt.Println(rotate([]string{"12","34"}, 2, true))
	fmt.Println(rotate([]string{"12","34"}, 3, true))
	fmt.Println(rotate([]string{"123","456","789"}, 1, true))
	fmt.Println(rotate([]string{"123","456","789"}, 2, true))
	fmt.Println(rotate([]string{"123","456","789"}, 3, true))
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
//
// ABC | JKL
// DEF | MNO
// GHI | PQR
// ---------
// STU | 123
// VWX | 456
// YZ0 | 789
func ExampleCombineGrids() {
	fmt.Println(combineGrids([][]string{{"12","34"}}))
	fmt.Println(combineGrids([][]string{{"12","34"},{"56","78"},{"AB","CD"},{"EF","GH"}}))
	fmt.Println(combineGrids([][]string{{"ABC","DEF","GHI"},{"JKL","MNO","PQR"},{"STU","VWX","YZ0"},{"123","456","789"}}))
	//Output:
	//[12 34]
	//[1256 3478 ABEF CDGH]
	//[ABCJKL DEFMNO GHIPQR STU123 VWX456 YZ0789]
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
//
// ABCDEF
// GHIJKL
// MNOPQR
// STUVWX
// YZ0123
// 456789
func ExampleDivideGrid() {
	fmt.Println(divideGrid([]string{"12","34"}))
	fmt.Println(divideGrid([]string{"123","456","789"}))
	fmt.Println(divideGrid([]string{"1234","5678","ABCD","EFGH"}))
	fmt.Println(divideGrid([]string{"ABCDEF", "GHIJKL", "MNOPQR", "STUVWX", "YZ0123", "456789"}))
	//Output:
	//[[12 34]]
	//[[123 456 789]]
	//[[12 56] [34 78] [AB EF] [CD GH]]
	//[[ABC GHI MNO] [DEF JKL PQR] [STU YZ0 456] [VWX 123 789]]
}