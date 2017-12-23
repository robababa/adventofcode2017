package quadratic

import (
	"math"
)

type solution interface {
	always() bool
	never() bool
	values() []int
}

type never struct {}

func (_ never) always() bool {
	return false
}

func (_ never) never() bool {
	return true
}

func (_ never) values() []int {
	var answer []int
	return answer
}

type always struct {}

func (_ always) always() bool {
	return true
}

func (_ always) never() bool {
	return false
}

// return an empty array. The caller needs to check always() first
func (_ always) values() []int {
	var answer []int
	return answer
}

type sometimes struct {
	solutionValues []int
}

func (_ sometimes) always() bool {
	return false
}

func (_ sometimes) never() bool {
	return false
}

func (s sometimes) values() []int {
	return s.solutionValues
}

// returns the positive integer solution to the equation bx + c = 0, if one exists
// if it is always true for all positive x, then it returns AlwaysZero
// if it is true for no positive integers x, then it returns NoSolutionValue
func linearPositiveIntegerSolution(b int, c int) solution {
	switch {
	// 0 = 0 is always true
	case b == 0 && c == 0: return always{}
		// c == 0 is always false when c != 0
	case b == 0 && c != 0: return never{}
		// bx = 0 for b != 0 is true iff x is 0
	case b != 0 && c == 0: return sometimes{solutionValues: []int{0}}
	default: {
		candidate := int(-1 * c / b)
		if candidate > 0 && b * candidate + c == 0 {
			return sometimes{solutionValues: []int{candidate}}
		} else {
			return never{}
		}
	}
	}
}

// returns the positive integer solutions to the quadratic equation
// ax^2 + bx + c = 0
// if there are two solutions, this function returns both of them
// if there is only one positive integer solution, it returns that solution and NoSolutionValue
// if there are no positive integer solutions, the function returns NoSolutionValue, NoSolutionValue
func quadraticPositiveIntegerSolutions(a int, b int, c int) solution {
	// first, get some edge cases out of the way when a = 0
	if a == 0 {
		return linearPositiveIntegerSolution(b, c)
	}

	// at this point, we have a real quadratic equation
	discriminant := b * b - 4 * a * c
	// imaginary solutions don't work here
	if discriminant < 0 {
		return never{}
	}

	// the discriminant is nonnegative.  However, if its root is not an integer, we have no solutions
	discriminantRoot := int(math.Sqrt(float64(discriminant)))
	if discriminantRoot * discriminantRoot != discriminant {
		return never{}
	}

	// put our potential solutions in ascending order
	couldBe := [2]int{
		((-1 * b) - discriminantRoot) / (2 * a),
			((-1 * b) + discriminantRoot) / (2 * a),
			}
	var reallyWorks []int

	// if the possible solutions don't actually work, set them to NoSolutionValue
	if couldBe[0] >= 0 && a * couldBe[0] * couldBe[0] + b * couldBe[0] + c == 0 {
		reallyWorks = append(reallyWorks, couldBe[0])
	}

	if couldBe[1] >= 0 && couldBe[1] != couldBe[0] && a * couldBe[1] * couldBe[1] + b * couldBe[1] + c == 0 {
		reallyWorks = append(reallyWorks, couldBe[1])
	}

	if len(reallyWorks) == 0 {
		return never{}
	}

	// reallyWorks has one or two elements, so there are some solution values
	return sometimes{reallyWorks}
}

