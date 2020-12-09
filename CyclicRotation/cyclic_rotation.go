//
// Title  : cyclic_rotation
// Author : Richie Varghese
// Date   : 08/12/2020
//
// Cyclic Rotation : An array A consisting of N integers is given. Rotation of the array means that each element is shifted right by one index,
//                   and the last element of the array is moved to the first place.
//                   For example, the rotation of array A = [3, 8, 9, 7, 6] is [6, 3, 8, 9, 7]
//                   (elements are shifted right by one index and 6 is moved to the first place).
//
// Functions : cyclicRotation ---> rotates the input array to right by n positions
//
// Result    : https://app.codility.com/demo/results/trainingQPF4JW-SH5/
//

package main

import "fmt"

//
// cyclicRotation function rotates an input array to its right by n positions
//
// INPUT  : inputArray int;
// RETURN : rotatedArray int
//
func cyclicRotation(inputArr []int, rotation int) []int {

	var updatedArr = make([]int, len(inputArr))

	// iterate on each entry in input array
	for origIndex := range inputArr {

		// set next index as current index + rotations in input
		// eg: if origIndex = 0, rotation = 3
		//        rotatedIndex = 0 + 3 = 3
		rotatedIndex := origIndex + rotation
		fmt.Println(rotatedIndex)

		// if next index is greater than the length of array, reduce the value to accomodate in array length
		if rotatedIndex >= len(inputArr) {

			// set next index as (origIndex + rotation) - length of arr
			// eg : if origIndex = 4, rotation = 3,  len of input arr = 6
			//         rotatedIndex = 4 + 3 = 7, out of bounds =>
			//                     ( 4 + 3 ) - 6 = 7 - 6 = 1
			rotatedIndex = (origIndex + rotation) - len(inputArr)

			// if rotated index  is still greater than input array, keep reducing len(array) until its smaller than input arr len
			// eg : if origIndex = 1, rotation = 3, len of input = 2
			//         rotatedIndex = 1 + 3 = 4, out of bounds =>
			//                      = ( 1 + 3 ) - 2 = 4 - 2 = 2, out of bounds =>
			//                      = 2 (prev result) - 2 (len of array) = 0, done
			//
			for rotatedIndex >= len(inputArr) {

				// reduce previous result with len of input
				rotatedIndex -= len(inputArr)
			}

		}

		// set the vaulue at origIndex of input array to next index of uptdaed Arr
		updatedArr[rotatedIndex] = inputArr[origIndex]
	}

	// return the updated array
	return updatedArr
}

// main function for performing cyclic rotation
func main() {

	var (
		rotation = 4
		inputArr = []int{1, 2}
	)

	// call cyclic rotation func for rotating the input array by rotation times
	result := cyclicRotation(inputArr, rotation)
	fmt.Println(result)

}
