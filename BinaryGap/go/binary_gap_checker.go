//
// Title  : binary_gap_checker
// Author : Richie Varghese
// Date   : 06/12/2020
//
// Binary Gap : A binary gap within a positive integer N is any maximal sequence of consecutive zeros that is surrounded
//              by ones at both ends in the binary representation of N.
//              N = 1041 the function should return 5, because N has binary representation 10000010001 and so its longest binary gap is of length 5.
//              N = 32 the function should return 0, because N has binary representation '100000' and thus no binary gaps
//
// Functions : binaryGapChecker(int) int ---> for returning the longest gap in a binary formatted integer (longest number of 0's)
//             main() ---> gets user input and calls binary gap checker func for result
//

package main

import (
	"fmt"
	"strconv"
)

//
// Solution function checks the number of 0's in the binary value of N and returns the longest number of 0's present for N
//
// INPUT  : input int;
// RETURN : longestGap int
//
func binaryGapChecker(input int) int {

	// longestGap is the return value expected for N after ops
	var longestGap, localMax int

	// convert the input number to its binary equivalent
	binaryVal := strconv.FormatInt(int64(input), 2)

	// iterate on each char from binary value
	for _, binElement := range binaryVal {

		// if current char is 0, add to local maximum as length of gap | else if char is 1, store local max to longest gap and reset localmax
		if string(binElement) == strconv.Itoa(0) {

			// add 1 to length of local maximum since 0 is found
			localMax++

		} else if string(binElement) == strconv.Itoa(1) {

			// check if local maximum is greater than longest gap
			if localMax > longestGap {

				// set longest gap as current local maximum
				longestGap = localMax
			}

			// reset local maximum
			localMax = 0

		}
	}

	// return the longest gap value for N
	return longestGap
}

// Main function for binary gap checker
func main() {

	var input int

	// get the input number from user for ops
	fmt.Printf("Insert value for number N : ")
	fmt.Scanf("%d\n", &input)

	// get the longest gap for an input integer
	result := binaryGapChecker(input)
	fmt.Printf("Longest gap for %d (%s) is %d\n", input, strconv.FormatInt(int64(input), 2), result)

}
