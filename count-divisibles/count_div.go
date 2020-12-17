//
// Title  : Count Divisibles
// Author : Richie Varghese
// Date   : 17/12/2020
//
// count_divisibles : given three integers A, B and K, returns the number of integers within the range [A..B] that are divisible by K, i.e.:
//                    { i : A ≤ i ≤ B, i mod K = 0 }
//                    For example, for A = 6, B = 11 and K = 2, your function should return 3, because there are three numbers divisible by 2 within the range [6..11], namely 6, 8 and 10.
//
// Functions : CountDiv ---> returns the number of divisibles by K for numbers within range of A and B
//
// Result    : https://app.codility.com/demo/results/trainingNFCXB9-FSU/
//
package countdiv

import "math"

//
// CountDiv function retuns the number of divisibles by K for numbers within range of A and B
//
// INPUT  : A, B, K int;
// RETURN : count int
//
func CountDiv(A, B, K int) int {

	if A%K == 0 {
		return int(math.Floor(float64(B/K))) - int(math.Floor(float64(A/K))) + 1
	} else {
		return int(math.Floor(float64(B/K))) - int(math.Floor(float64(A/K)))
	}
}
