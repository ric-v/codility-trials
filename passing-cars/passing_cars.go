//
// Title  : Passing Cars
// Author : Richie Varghese
// Date   : 20/12/2020
//
// passing cars : A non-empty array A consisting of N integers is given. The consecutive elements of array A represent consecutive cars on a road.
//                Array A contains only 0s and/or 1s:
//                0 represents a car traveling east,
//                1 represents a car traveling west.
//                The goal is to count passing cars. We say that a pair of cars (P, Q), where 0 ≤ P < Q < N, is passing when P is traveling to the east and Q is traveling to the west.
//                For example, consider array A such that:
//                  A[0] = 0
//                  A[1] = 1
//                  A[2] = 0
//                  A[3] = 1
//                  A[4] = 1
//                We have five pairs of passing cars: (0, 1), (0, 3), (0, 4), (2, 3), (2, 4).
//
// Functions : PassingCars ---> returns number of passing car pairs from input array
//
// Results   : https://app.codility.com/demo/results/trainingX2ZYSZ-ERZ/
//
package passingcars

import "fmt"

//
// PassingCars function returns the number of passing car pairs from input array
//
// INPUT  : input []int;
// RETURN : numOfPairs int
//
func PassingCars(input []int) int {

	var numOfPairs int
	var ref = 0

	for index := range input {

		if input[index] == ref {

			var innerIndex = index

			for innerIndex < len(input) {

				fmt.Println(index, input[index], innerIndex, input[innerIndex])
				if ref != input[innerIndex] {
					numOfPairs++
					if numOfPairs == 1000000000 {
						numOfPairs = -1
						break
					}
					fmt.Println(numOfPairs)
				}
				innerIndex++
			}
			if numOfPairs == -1 {
				break
			}
		}
	}

	return numOfPairs
}
