/*
		You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order,
		and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.
		You may assume the two numbers do not contain any leading zero, except the number 0 itself.

		1. get integers
			- iterate over list adding digit to sum after applying multiplier to digit
			- pass first ListNode to my iterator function along with function:
				
		2. get digits
			- sum % 10 / 1 gives ones place
				- subtract ones place and make that new sum; multiply to mod and divisor by 10
			- sum % 100 / 10 gives hundreds place
				- subtract ones place and make that new sum; multiply to mod and divisor by 10
			- sum % 1000 / 100 gives thousands place
				- etc
			- do until mod is greater than sum, then do that one last time

		Definition for singly-linked list:

  		type ListNode struct {
      		Val int
      		Next *ListNode
 		}


		Example 1:

		  2  ->  4  ->  3
		+ 5  ->  6  ->  4
		_________________
		  7  ->  0  ->  8


		Input: l1 = [2,4,3], l2 = [5,6,4]
		Output: [7,0,8]
		Explanation: 342 + 465 = 807



		Example 2:

		  0
		+ 0
		___
		  0


		Input: l1 = [0], l2 = [0]
		Output: [0]


		Input
		[1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1]
		[5,6,4]

		Expected
		[6,6,4,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1]
*/
package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	l1 := sliceToList([]int64{1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1})
	l2 := sliceToList([]int64{5,6,4})
	sum := addTwoNumbers(l1, l2)
	listEach(sum, func(l ListNode) {
		fmt.Printf("%d ", l.Val)
	})
	fmt.Println()
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var sum1 int64 = 0
	var sum2 int64 = 0
	var multiplier int64 = 1
    listEach(l1, func(l ListNode) {
		val := (int64(l.Val) * multiplier)
		sum1 += val
		multiplier *= 10
	})
	multiplier = 1
	listEach(l2, func(l ListNode) {
		val := (int64(l.Val) * multiplier)
		sum2 += val
		multiplier *= 10
	})

	sum := sum1 + sum2
	digits := getDigits(sum)
	return sliceToList(digits)
}

/*
	takes a pointer to a ListNode and treats it as the head of a linked-list
	takes a function which receives a ListNode, and performs its action on the ListNode and all subsequent linked-nodes
*/
func listEach(l *ListNode, fn func(ListNode)) int {
	node  := *l
	count := 0
	for {
		fn(node)
		count++
		if (node.Next != nil) {
			node = *node.Next
		} else {
			break
		}
	}
	return count
}

/*
	converts a slice of ints to a linked list, returning a pointer to the head ListNode
*/
func sliceToList(slice []int64) *ListNode {
	head := ListNode{int(slice[0]), nil}
	var curr *ListNode = &head

	for i := 1; i < len(slice); i++ {
		newNode := ListNode{int(slice[i]), nil}
		curr.Next = &newNode
		curr = &newNode
	}

	return &head
}

func getDigits(i int64) []int64 {
	digits := []int64{}
	var mod int64 = 10
	var div int64 = 1
	for {
		digit := (i % mod) / div
		i = i - (digit * div)
		digits = append(digits, digit)
		if (mod > i) {
			break
		}
		mod *= 10
		div *= 10
	}
	return digits
}