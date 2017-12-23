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

	//Output:
	//2 , -6 always(), never(), values() are false false [3]
	//2 , -5 always(), never(), values() are false true []
	//2 , 6 always(), never(), values() are false true []
	//2 , 0 always(), never(), values() are false false [0]
	//0 , 0 always(), never(), values() are true false []

}

//func ExampleQuadraticPositiveIntegerSolutions() {
//
//}
