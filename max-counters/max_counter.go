//
// Title  : max_counter
// Author : Richie Varghese
// Date   : 15/12/2020
//
// MaxCounter : You are given N counters, initially set to 0, and you have two possible operations on them:
//              increase(X) − counter X is increased by 1,
//              max counter − all counters are set to the maximum value of any counter.
//              A non-empty array A of M integers is given. This array represents consecutive operations:
//              if A[K] = X, such that 1 ≤ X ≤ N, then operation K is increase(X),
//              if A[K] = N + 1 then operation K is max counter.
//              For example, given:
//                A[0] = 3
//                A[1] = 4
//                A[2] = 4
//                A[3] = 6
//                A[4] = 1
//                A[5] = 4
//                A[6] = 4
//              the function should return [3, 2, 2, 4, 2], as explained above.
//
// Functions : MaxCounter ---> returns the array of max counters based on the instruction set give
//
package maxcounters

//
// MaxCounter function returns the array of max counters based on instruction set
//
// INPUT  : input []int, N int;
// RETURN : result []int
//
func MaxCounter(input []int, N int) []int {

	var (
		localMax int
		result   = make([]int, N)
	)

	for _, elem := range input {

		if elem < N+1 {
			result[elem-1]++
			localMax = result[elem-1]
		} else {

			var index int

			for index < len(result) {
				result[index] = localMax
				index++
			}
		}
	}

	return result
}
