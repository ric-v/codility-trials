//
// Title  : odd_occurrences_in_array
// Author : Richie Varghese
// Date   : 09/12/2020
//
// OddOccurrencesInArray : A non-empty array A consisting of N integers is given.
//                         The array contains an odd number of elements, and each element of the array can be paired with another element
//                         that has the same value, except for one element that is left unpaired.
//                         For example, given array A such that:
//                         A[0] = 9  A[1] = 3  A[2] = 9
//                         A[3] = 3  A[4] = 9  A[5] = 7
//                         A[6] = 9
//                         the function should return 7, as explained in the example above.
//
// Functions : oddOccurrencesInArray --->
//
// Result    : https://app.codility.com/demo/results/trainingZQ5HYK-NRH/
//

package oddoccurrencesinarray

//
// oddOccurrencesInArray function finds the value with odd occurrence in an input array.
//
// INPUT  : inputArr []int;
// RETURN : result int
//
func OddOccurrencesInArray(inputArr []int) int {

	const limit = 2

	var (
		result             int
		occurrenceCheckMap = make(map[int]int)
	)

	// iterate on each element in input array, for each occurance update occurance map by 1
	for _, entry := range inputArr {

		// update occurance check map by 1
		occurrenceCheckMap[entry]++

		// if occurrenceCheckMap occured limit times, ignore the entry
		if occurrenceCheckMap[entry] == limit {
			delete(occurrenceCheckMap, entry)
		}
	}

	// get the only entry from occurrence checker map
	for entry := range occurrenceCheckMap {
		result = entry
	}

	return result
}
