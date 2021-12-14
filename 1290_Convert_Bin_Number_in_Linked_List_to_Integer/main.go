/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }


	To avoid having to traverse list to get count/most-sginificant bit
		start at beginning, and keep track of the power by multiplying sum by 2 
		for each node

 */
package main

type ListNode struct {
	Val int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	curr := head
	sum := 0

	for {
		sum *= 2
		sum += curr.Val
		if curr.Next == nil {
			break
		}
		curr = curr.Next
	}

	return sum
}