//
// Title  : missing_element
// Author : Richie Varghese
// Date   : 11/12/2020
//
// MissingElement : An array A consisting of N different integers is given. The array contains integers in the range [1..(N + 1)], which means that exactly one element is missing.
//                  For example, given array A such that:
//                  A[0] = 2
//                  A[1] = 3
//                  A[2] = 1
//                  A[3] = 5
//                  the function should return 4, as it is the missing element.
//
// Functions : MissingElement ---> returns the missing element from an array with N+1 concurrent integers with 1 among them missing
//
// Result    : https://app.codility.com/demo/results/trainingSREVS9-8CU/
//
package missingelement

import "sort"

//
// MissingElement function returns the missing element in a continuous array of int
//
// INPUT  : intput []int;
// RETURN : missingElem int
//
func MissingElement(input []int) int {

	var missingElem = 1

	// sort the slice in ascending order
	sort.Ints(input)

	// iterate on each input element
	for _, element := range input {
		if element != missingElem {
			break
		} else {
			missingElem++
		}
	}

	return missingElem
}
