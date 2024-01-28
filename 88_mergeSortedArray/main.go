package main

import "fmt"

/*
	You are given two integer arrays nums1 and nums2, sorted in increasing order, and two
		integers m and n, representing the number of elements in nums1 and nums2 respectively.

	Merge nums1 and nums2 into a single array sorted in increasing order.

	The final sorted array should not be returned by the function, but instead be stored
		inside the array nums1. To accommodate this, nums1 has a length of m + n, where
		the first m elements denote the elements that should be merged, and the last n
		elements are set to 0 and should be ignored. nums2 has a length of n.

	Example 1:

		Input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
		Output: [1,2,2,3,5,6]
		Explanation: The arrays we are merging are [1,2,3] and [2,5,6].
		The result of the merge is [1,2,2,3,5,6] with the underlined elements coming from nums1.

	Example 2:

		Input: nums1 = [1], m = 1, nums2 = [], n = 0
		Output: [1]
		Explanation: The arrays we are merging are [1] and [].
		The result of the merge is [1].

	given two sorted (increasing) arrays of ints, add the ints from the second array into the first array, maintaining the order

.
	1  1  2  2  3  4
	   ^

	1  2  4
.



	1  1  2  2  3  4



  .
	1  3  4   5   23   67  70  80  81
		^

	1   3   5  67  70  80
	.


	1  3  4  5  23  67  70  80  81

	set 3 pointers
		a, b, end

	a is set to greatest in nums1, just before the filling 0s (m - 1)
	b is the greatest in nums2, so the end of nums2 (n - 1)
	end is the end of nums2 (m+n - 1)

	as long as end is greater than or equal to 0
		(ensure a and b are still valid indexes)
		find greater of a and b
			if a
				set end to a
				decrement end
				decrement a pointer
			if b
				set end to b
				decrement end
				decrement b pointer
			if equal
				follow b case
	
	return  nums1

.
	2 4 3 4 5 6
      e
	1 3 5
	  .


[2 2 3 4 5 6]



*/

func merge(nums1 []int, m int, nums2 []int, n int)  {
	a, b, e := m-1, n-1, (m+n) - 1
    for ; a >= 0 && b >= 0; {
		if nums1[a] > nums2[b] {
			nums1[e] = nums1[a]
			a--
		} else {
			nums1[e] = nums2[b]
			b--
		}
		e--
	}
	if a < 0 {
		for ;b >= 0; {
			nums1[b] = nums2[b]
			b--
		}
	}
}


func main() {
	n1 := []int{3,10,11,18,29,700,0,0,0,0,0}
	m  := 6
	n2 := []int{1,47,289,300,700}
	n := 5

	fmt.Println(n1, n2)

	merge(n1, m, n2, n)

	fmt.Println(n1)
}