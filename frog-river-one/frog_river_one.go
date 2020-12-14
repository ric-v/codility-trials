//
// Title  : frog_river_one
// Author : Richie Varghese
// Date   : 14/12/2020
//
// Frog River One : A small frog wants to get to the other side of a river. The frog is initially located on one bank of the river (position 0) and wants to
//                  get to the opposite bank (position X+1). Leaves fall from a tree onto the surface of the river.
//                  You are given an array A consisting of N integers representing the falling leaves. A[K] represents the position where one leaf falls at time K, measured in seconds.
//                  The goal is to find the earliest time when the frog can jump to the other side of the river. The frog can cross only when leaves appear at every position across the
//                  river from 1 to X (that is, we want to find the earliest moment when all the positions from 1 to X are covered by leaves).
//                  You may assume that the speed of the current in the river is negligibly small, i.e. the leaves do not change their positions once they fall in the river.
//                  For example, given X = 5 and array A such that:
//                    A[0] = 1
//                    A[1] = 3
//                    A[2] = 1
//                    A[3] = 4
//                    A[4] = 2
//                    A[5] = 3
//                    A[6] = 5
//                    A[7] = 4
//                  the function should return 6, as explained above.
//
// Functions : FrogRiverOne ---> returns the smallest time at which frog can make the jump to cross the river
//
// Result    : https://app.codility.com/demo/results/training7DZANE-VWY/
//

package frogriverone

import "fmt"

//
// FrogRiverOne function returns the smallest time at which frog can make the jump to cross the river
//
// INPUT  : input []int, x int;
// RETURN : smallestTime int
//
func FrogRiverOne(input []int, x int) int {

	var (

		// initialize smallest time as -1 or non existant
		index, smallestTime = 1, -1
		leavesMissing       = make(map[int]bool)
	)

	// iterate and set missing leaves to map for tracking a new leaf when it falls
	for index <= x {
		leavesMissing[index] = true
		index++
	}

	// iterate on each timeslot and check for the falling leaf
	for timeNow, elem := range input {

		// if the falling leaf is in the missing leaves, remove the entry as it is now found in river
		if _, ok := leavesMissing[elem]; ok {
			fmt.Println(leavesMissing)
			delete(leavesMissing, elem)
		}

		// if all missing leaves are discovered, break from loop and set current index as smallest time
		if len(leavesMissing) == 0 {
			smallestTime = timeNow
			break
		}
	}

	return smallestTime
}
