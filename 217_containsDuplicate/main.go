package main

import "fmt"
/*
	Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.

 

	Example 1:

	Input: nums = [1,2,3,1]
	Output: true

	Example 2:

	Input: nums = [1,2,3,4]
	Output: false

	Example 3:

	Input: nums = [1,1,1,3,3,4,3,2,4,2]
	Output: true


	approach 1:

	two-pointers, tortoise-hare
	tortoise visits each while hare visits each greater index to check for duplicates
	time complexity: O(n^2) - visiting every element and for each visit we visit all other elelments (greater than current anyway)
	space: O(1) - the only memory we use is to define the two pointers

	approach 2:

	map
	visit each element only once, using a map to keep track of which values have been seen
	time complexity: O(n) - visit each element once
	space complexity: O(n) - map grows with input

*/

func main() {
	fmt.Println(containsDuplicate([]int{}) == false)
	fmt.Println(containsDuplicate([]int{1}) == false)
	fmt.Println(containsDuplicate([]int{1,2,3}) == false)
	fmt.Println(containsDuplicate([]int{1,1,2,3}))
	fmt.Println(containsDuplicate([]int{1,2,2,3}))
	fmt.Println(containsDuplicate([]int{1,2,3,3}))
	fmt.Println(containsDuplicate([]int{1,2,3,1}))
	fmt.Println(containsDuplicate([]int{2,2}))
	fmt.Println(containsDuplicate([]int{1,1,1,1,1,1,1,1,1,1}))
}

func containsDuplicate(nums []int) bool {

	length := len(nums)
	seen := make(map[int]bool, length)

	for i := 0; i < length; i++ {

		if haveSeen, _ := seen[nums[i]]; haveSeen {
			return true;
		}

		seen[nums[i]] = true

	}

	return false;
}