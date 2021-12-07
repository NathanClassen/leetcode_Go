/*
		Given an array of integers `nums`` and an integer `target``, return indices of the two numbers such that they add up to `target`.
        You may assume that each input would have exactly one solution, and you may not use the same element twice.
        You can return the answer in any order.
		_______________________________________

		- Two pointers `l` `r`, start at 0 and 1 respectively
		- increment r, checking each index against l to see if they sum to target, until it reaches nums.length - 1
		- at which point l is incremented, r is set to l + 1, and step two repeats
		- if l equals nums.length - 2, we have reached the max that either pointer can be, and because the problem notes that
			there is always a solution, we have found the solution and may return l and r
		_____________________________________________________________________________________________________________________

		One iteration over `nums`
		For each num
			- check map with key: (target - nums[i])
				- if entry is found the (nums[i]) + map[target - nums[i]] == target; return []int{map[target - nums[i]], i}
			- save entry to map; key is (nums[i]) & value is index

*/

package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2,7,-11,15}, -9))
}

/* 
		Time  complexity: O(n)
		Space complextiy: O(n)
*/

func twoSum(nums []int, target int) []int {
	var complements = map[int]int{}
	var res []int

	for i := 0; i < len(nums); i++ {
		if _, flag := complements[(target - nums[i])]; flag {
			res = []int{(complements[(target - nums[i])]), i}
			break
		}
		complements[(nums[i])] = i
	}
	return res
}

/*
		Time  complexity: O(n^2)
		Space complexity: O(1)
*/

func twoSumSlower(nums []int, target int) []int {
	var l int = 0
	var r int

	for ; l <= (len(nums) - 2); l += 1 {
		for r = l + 1; r <= (len(nums) - 1); r += 1 {
			if nums[l] + nums[r] == target {
				return []int{l, r}
			}
		}
	}

	return []int{l, r}
}

/*



Example 1:

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Output: Because nums[0] + nums[1] == 9, we return [0, 1].

Example 2:

Input: nums = [3,2,4], target = 6
Output: [1,2]

Example 3:

Input: nums = [3,3], target = 6
Output: [0,1]

*/