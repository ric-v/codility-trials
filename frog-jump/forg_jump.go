//
// Title  : frog_jump
// Author : Richie Varghese
// Date   : 11/12/2020
//
// Frog Jump : A small frog wants to get to the other side of the road. The frog is currently located at position X and wants to get to a
//             position greater than or equal to Y. The small frog always jumps a fixed distance, D. Count the minimal number of jumps that the
//             small frog must perform to reach its target.
//
//             For example, given:
//                X = 10
//                Y = 85
//                D = 30
//              the function should return 3, because the frog will be positioned as follows:
//                after the first jump, at position 10 + 30 = 40
//                after the second jump, at position 10 + 30 + 30 = 70
//                after the third jump, at position 10 + 30 + 30 + 30 = 100
//
// Functions : FrogJump ---> returns number of jumps needed to cross the road where x + jump > y
//
// Result    : https://app.codility.com/demo/results/trainingERYFWG-SG2/
//

package frogjmp

import (
	"math"
)

//
// FrogJump function returns the number of jumps needed to cross the road where x + jump > y
//
// INPUT  : X, Y, jumpDist int;
// RETURN : numOfJumps int
//
func FrogJump(x, y, jumpDist int) int {

	return int(
		math.Ceil(
			(float64(y) - float64(x)) / float64(jumpDist),
		),
	)
}
