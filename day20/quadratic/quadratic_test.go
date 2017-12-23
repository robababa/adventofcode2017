package quadratic

import "fmt"

func ExampleLinearPositiveIntegerSolution() {
	// 2x - 6 = 0 : x = 3 is a solution
	s1 := linearPositiveIntegerSolution(2, -6)
	fmt.Println("(2, -6) always(), never(), values() are", s1.always(), s1.never(), s1.values())

	// 2x - 5 = 0 : x = 2.5 is not an integer, so there are no solutions
	s2 := linearPositiveIntegerSolution(2, -5)
	fmt.Println("(2, -5) always(), never(), values() are", s2.always(), s2.never(), s2.values())

	// 2x + 6 = 0 : x = -3 is not positive, so there are no solutions
	s3 := linearPositiveIntegerSolution(2, 6)
	fmt.Println("(2,  6) always(), never(), values() are", s3.always(), s3.never(), s3.values())

	// 2x = 0 : x = 0 is a solution
	s4 := linearPositiveIntegerSolution(2, 0)
	fmt.Println("(2,  0) always(), never(), values() are", s4.always(), s4.never(), s4.values())

	//Output:
	//(2, -6) always(), never(), values() are false false [3]
	//(2, -5) always(), never(), values() are false true []
	//(2,  6) always(), never(), values() are false true []
	//(2,  0) always(), never(), values() are false false [0]
}

//func ExampleQuadraticPositiveIntegerSolutions() {
//
//}
