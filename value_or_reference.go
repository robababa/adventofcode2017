package main

import "fmt"

func main() {
	fmt.Println("Simple program to show whether variables are passed by value or by reference")
	i := 1
	s := "Original string"
	a := []int{1, 2, 3}
	m := make(map[int]string)
	m[10] = "Original value for index 10"
	fmt.Println("Original values:")
	printValues(i, s, a, m)
	fmt.Println()
	mutateValues(i, s, a, m)
	fmt.Println()
	fmt.Println("Values in main function after mutating them as function arguments:")
	printValues(i, s, a, m)

}

func printValues(i int, s string, a []int, m map[int]string) {
	fmt.Println("Integer is", i)
	fmt.Println("String is", s)
	fmt.Println("Array is", a)
	fmt.Println("Map is", m)
}

func mutateValues(i int, s string, a []int, m map[int]string) {
	i = 79
	s = "Mutated string"
	a[0] = 23
	m[19] = "Mutated value for 19"
	fmt.Println("Values in mutateValues() function after mutating them:")
	printValues(i, s, a, m)
}
