/*
	Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.
	The overall run time complexity should be O(log (m+n)).

	Example 1:

	Input: nums1 = [1,3], nums2 = [2]
	Output: 2.00000
	Explanation: merged array = [1,2,3] and median is 2.

	Example 2:

	Input: nums1 = [1,2], nums2 = [3,4]
	Output: 2.50000
	Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.

	- Sorted arrays (can use binary operations)
	- Median will equal the middle digit once the two arrays are merged
		- if there is one digit in the middle, this will of course be the median
		- if there are two digits in the middle, the median will be the two digits averaged

	- Will need two merge the arrays
		- that they are sorted will come in handy here for keeping the time complexity down

	High level overview
		- merge two arrays
		- if length of merged array is odd
			return middle digit
		- else
			return middle two digits averaged
	If m or n is 0, return the median of the other array
	? What about if both ar 0?

	To merge set pointer at the beginning of each array
		if l < r
			add l to result set
		if r < l
			add r to result set
		else they are equal
			add l to set

		increment pointer of the digit that was added

		if (l || r) moves past (m || n) respectively
			end loop, and addd remaining digits of the other array to the set

	Constraints:

    nums1.length == m
    nums2.length == n
    0 <= m <= 1000
    0 <= n <= 1000
    1 <= m + n <= 2000
    -10^6 <= nums1[i], nums2[i] <= 10^6


	Need to modify to get to O(logn) time


*/

package main

import "fmt"

func main() {
	findMedianSortedArrays([]int{1, 2}, []int{3, 4})
}

func findMedianSortedArrays(m []int, n []int) float32 {
	res := merge(m, n)
	fmt.Println(res)
	if len(res)%2 == 0 {
		i := len(res) / 2
		x := i - 1
		return (float32(res[i]) + float32(res[x])) / 2.0
	} else {
		return float32(res[len(res)/2])
	}
}

func merge(m []int, n []int) []int {
	if len(m) == 0 {
		return n
	} else if len(n) == 0 {
		return m
	}

	var res []int
	var l int
	var r int

	for l, r = 0, 0; l < len(m) && r < len(n); {
		if m[l] < n[r] {
			res = append(res, m[l])
			l++
		} else if n[r] < m[l] {
			res = append(res, n[r])
			r++
		} else {
			res = append(res, m[l])
			l++
		}
	}

	if l >= len(m) && r < len(n) {
		for ; r < len(n); r++ {
			res = append(res, n[r])
		}
	} else if r >= len(n) && l < len(m) {
		for ; l < len(m); l++ {
			res = append(res, m[l])
		}
	}

	return res
}
