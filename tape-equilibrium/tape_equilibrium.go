//
// Title  : tape_equilibrium
// Author : Richie Varghese
// Date   : 12/12/2020
//
// Tape Equilibrium : A non-empty array A consisting of N integers is given. Array A represents numbers on a tape.
//                    Any integer P, such that 0 < P < N, splits this tape into two non-empty parts: A[0], A[1], ..., A[P − 1] and A[P], A[P + 1], ..., A[N − 1].
//                    The difference between the two parts is the value of: |(A[0] + A[1] + ... + A[P − 1]) − (A[P] + A[P + 1] + ... + A[N − 1])
//                    For example, consider array A such that:
//                      A[0] = 3
//                      A[1] = 1
//                      A[2] = 2
//                      A[3] = 4
//                      A[4] = 3
//                    We can split this tape in four places:
//                    P = 1, difference = |3 − 10| = 7
//                    P = 2, difference = |4 − 9| = 5
//                    P = 3, difference = |6 − 7| = 1
//                    P = 4, difference = |10 − 3| = 7
//                    the function should return 1
//
// Functions : TapeEquilibrium ---> returns the smallest difference between tape equilibrium for an int array
//
package TapeEquilibrium

import (
	"math"
)

//
// TapeEquilibrium function returns the tape equilibrium for input array
//
// INPUT  : input []int;
// RETURN : smallest diff int
//
func TapeEquilibrium(input []int) int {

	arraySum := 0
	currentMin := 1<<32 - 1

	for _, value := range input {
		arraySum += value
	}
	lhs := input[0]
	rhs := arraySum - lhs

	for i := 1; i < len(input); i++ {
		diff := int(math.Abs(float64(lhs) - float64(rhs)))

		if diff < currentMin {
			currentMin = diff
		}
		lhs += input[i]
		rhs -= input[i]
	}

	return currentMin
}
