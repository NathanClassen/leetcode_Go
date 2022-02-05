/*
	Given a signed 32-bit integer x, return x with its digits reversed.
	If reversing x causes the value to go outside the signed 32-bit integer range [-231, 231 - 1], then return 0.

	Assume the environment does not allow you to store 64-bit integers (signed or unsigned).

	Example 1:
	Input: x = 1234
	Output: 4321

- 4
123
- 3
12
- 2
1


	to get hundreds place digit (h) -> x % 10
	to get tens place digit (t) -> ((x - h) % 100) / 10
	to get ones place digit (o) -> (x - (h + t)) / 100


	- can get results highest power by modulous 10
	- successive digits are found by modulous(lastMod * 10) / divisor (starts at 1; also multiplied by 10 each time)
		- how to aboid overflow
			- can actually just use mod and divisions of 10 to obtain digits, and break number down
				- digit = x % 10
				- then x = x-digit
				- then x = x/10

			- this eventually lands at 1 digit (x will be < 10) and thats the final digit, or the ones place

		- how to apply the correct multiplier though...
			4
			40 + 3
			43
			430 + 2
			432
			4320 + 1
		- the addition of each digit shall proceed as follows (assuming sum starts at 0)
		- multiply sum by 10, and add digit


		algorithm

		- set `sum` to 0
		- set sign to (if x < 0 -1, else 1)
		- while ??
			- get digit
			- reassign x
			- multiply sum by 10
			- add digit


	Example 2:
	Input: x = -123
	Output: -321

	Example 3:
	Input: x = 120
	Output: 21

	Constraints:
		-231 <= x <= 231 - 1


*/

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(reverse(-11234))
}

func reverse(x int) int {
	sum := 0

	for {
		digit := x % 10
		x = (x - digit) / 10
		sum *= 10
		sum += digit
		if x == 0 {
			break
		}
	}

	if sum > math.MaxInt32 || sum < math.MinInt32 {
		return 0
	}

	return sum
}
