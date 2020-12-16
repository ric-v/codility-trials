//
// Title  : Permutation Checker
// Author : Richie Varghese
// Date   : 16/12/2020
//
// Permutation Checker : A non-empty array A consisting of N integers is given.
//                       A permutation is a sequence containing each element from 1 to N once, and only once.
//                       For example, array A such that:
//                           A[0] = 4
//                           A[1] = 1
//                           A[2] = 3
//                           A[3] = 2
//                       is a permutation, but array A such that:
//                           A[0] = 4
//                           A[1] = 1
//                           A[2] = 3
//                       is not a permutation, because value 2 is missing.
//                       The goal is to check whether array A is a permutation.
//
// Functions : PermChecker ---> returns if an array is a permutation or not
//
package permcheck

import (
	"sort"
)

//
// PermChecker function checks if an input array is permutation or not
//
// INPUT  : input []int;
// RETURN : result int
//
func PermChecker(input []int) int {

	var result, iter = 1, 1

	// sort the values in ascending order
	sort.Ints(input)

	// iterate on each element in input array
	for _, elem := range input {

		// if current entry matches iter, update iter by 1
		if elem == iter {
			iter++
		} else {
			result = 0
			break
		}
	}

	return result
}
