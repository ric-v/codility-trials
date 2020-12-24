//
// Title  : Distinct
// Author : Richie Varghese
// Date   : 24/12/2020
//
// distict :  given an array A consisting of N integers, returns the number of distinct values in array A.
//            For example, given array A consisting of six elements such that:
//            A[0] = 2    A[1] = 1    A[2] = 1
//            A[3] = 2    A[4] = 3    A[5] = 1
//            the function should return 3, because there are 3 distinct values appearing in array A, namely 1, 2 and 3.
//
// Functions : distinct ---> returns number of distinct elements in an array
//
// Result    : https://app.codility.com/demo/results/trainingNKFWWC-5VW/
//
package distinct

//
// Distinct function returns number of distinct elements in an array
//
// INPUT  : input []int;
// RETURN : distinctEntries int
//
func Distinct(input []int) int {

	var distinctEntries = make(map[int]bool)

	// iterate on each entry from input array
	for _, entry := range input {
		distinctEntries[entry] = true
	}

	return len(distinctEntries)
}
