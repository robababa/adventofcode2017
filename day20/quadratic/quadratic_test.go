package quadratic

import "fmt"

func linearHelper(b, c int) {
	sol := linearPositiveIntegerSolution(b, c)
	fmt.Println(b, ",", c, "always(), never(), values() are", sol.always(), sol.never(), sol.values())
}

func ExampleLinearPositiveIntegerSolution() {
	// 2x - 6 = 0 : x = 3 is a solution
	linearHelper(2, -6)

	// 2x - 5 = 0 : x = 2.5 is not an integer, so there are no solutions
	linearHelper(2, -5)

	// 2x + 6 = 0 : x = -3 is not positive, so there are no solutions
	linearHelper(2, 6)

	// 2x = 0 : x = 0 is a solution
	linearHelper(2, 0)

	// 0x = 0 : always is the solution
	linearHelper(0, 0)

	// Output:
	// 2 , -6 always(), never(), values() are false false [3]
	// 2 , -5 always(), never(), values() are false true []
	// 2 , 6 always(), never(), values() are false true []
	// 2 , 0 always(), never(), values() are false false [0]
	// 0 , 0 always(), never(), values() are true false []
}

func quadraticHelper(a, b, c int) {
	sol := quadraticPositiveIntegerSolutions(a, b, c)
	fmt.Println(a, ",", b, ",", c, "always(), never(), values() are", sol.always(), sol.never(), sol.values())
}

func ExampleQuadraticPositiveIntegerSolutions() {
	// 1*x^2 - 4x + 3 = 0 : x = 1, 3
	quadraticHelper(1, -4, 3)
	// x^2 - 4 = 0 : x = 2
	quadraticHelper(1, 0, -4)
	// x^2 - 4x + 4 = 0 : x = 2 is a double solution, but should only appear once
	quadraticHelper(1, -4, 4)
	// x^2 + 4x + 3 = 0 : no solutions
	quadraticHelper(1, 4, 3)

	fmt.Println()
	fmt.Println("a = 0:")
	// 0x^2 + x - 7 = 0 : x = 7
	quadraticHelper(0, 1, -7)

	fmt.Println()
	fmt.Println("a = b = 0, c nonzero")
	// 2 = 0 : no solutions
	quadraticHelper(0, 0, 2)

	fmt.Println()
	fmt.Println("a = b = c = 0")
	// 0 = 0 : always
	quadraticHelper(0, 0, 0)
	// Output:
	// 1 , -4 , 3 always(), never(), values() are false false [1 3]
	// 1 , 0 , -4 always(), never(), values() are false false [2]
	// 1 , -4 , 4 always(), never(), values() are false false [2]
	// 1 , 4 , 3 always(), never(), values() are false true []
	//
	// a = 0:
	// 0 , 1 , -7 always(), never(), values() are false false [7]
	//
	// a = b = 0, c nonzero
	// 0 , 0 , 2 always(), never(), values() are false true []
	//
	// a = b = c = 0
	// 0 , 0 , 0 always(), never(), values() are true false []
}

func ExampleCombineSolutions() {
	fmt.Println("combineSolutions(always{}, always{}).always() =",
		combineSolutions(always{}, always{}).always())
	fmt.Println("combineSolutions(always{}, never{}).never() =",
		combineSolutions(always{}, never{}).never())
	fmt.Println("combineSolutions(always{}, sometimes{solutionValues: []int{5, 4}}).values() =",
		combineSolutions(always{}, sometimes{solutionValues: []int{5, 4}}).values())
	fmt.Println("combineSolutions(never{}, sometimes{solutionValues: []int{5, 4}}).never() =",
		combineSolutions(never{}, sometimes{solutionValues: []int{5, 4}}).never())
	fmt.Println("combineSolutions(sometimes{solutionValues: []int{3, 9, 7}}, sometimes{solutionValues: []int{5, 7, 4}}).values() =",
		combineSolutions(sometimes{solutionValues: []int{3, 9, 7}}, sometimes{solutionValues: []int{5, 7, 4}}).values())
	// Output:
	// combineSolutions(always{}, always{}).always() = true
	// combineSolutions(always{}, never{}).never() = true
	// combineSolutions(always{}, sometimes{solutionValues: []int{5, 4}}).values() = [4]
	// combineSolutions(never{}, sometimes{solutionValues: []int{5, 4}}).never() = true
	// combineSolutions(sometimes{solutionValues: []int{3, 9, 7}}, sometimes{solutionValues: []int{5, 7, 4}}).values() = [7]
}