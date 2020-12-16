//
// Title  : missing_integer
// Author : Richie Varghese
// Date   : 16/12/2020
//
// Missing Integer : given an array A of N integers, returns the smallest positive integer (greater than 0) that does not occur in A.
//                     For example, given A = [1, 3, 6, 4, 1, 2], the function should return 5.
//                     Given A = [1, 2, 3], the function should return 4.
//                     Given A = [−1, −3], the function should return 1.
//
// Functions : MissingInteger ---> returns the smallest integer that does not occur in A
//
// Result    : https://app.codility.com/demo/results/trainingYNQY6Y-2SY/
//
package missingint

import (
	"sort"
)

//
// MissingInteger function returns smallest int that does not occur in A
//
// INPUT  : input []int;
// RETURN : result int
//
func MissingInteger(input []int) int {

	// init result as 1 (the least positive possible int)
	var result = 1

	// sort the input array in ascending order
	sort.Ints(input)

	// iterate on each entry from input array
	for _, elem := range input {

		// if current entry is result, update result by 1, else do nothing
		if elem == result {
			result++
		}

		// if result +1 returns value less than current element, current result is the smallest possible integer, break the loop
		if result+1 < elem {
			break
		}
	}

	return result
}
