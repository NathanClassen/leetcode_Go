/*
	Given the head of a singly linked list, reverse the list, and return the reversed list.

	Input: head = [1,2,3,4,5]
	Output: [5,4,3,2,1]

	Input: head = [1,2]
	Output: [2,1]

	Input: head = []
	Output: []
*/
/*

             p  c
d < 1 < 2  < 3  4
    h
	to reverse a list, must turn all pointers (next) around
	to traverse list, take next node as long as next is not nil

	if empty node or single node with no next node return that node

		start at beginning of list, and point first node at a dummy node and maintain a ref to this first node
		traverse list setting the next node to point at the current node, as long as there is a next node
		finally, set first node ref.next to nil

	algo

	create dummy node, set it as the prev ref
	create ref to first node (head ref), and set it as the (curr_node ref)
	for as c.next not nil
		get ref to next node
		set c.next to p
		set p to be c
		set next to c

	set c.next to p
	set h.next to nil
	return c
*/

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dum := ListNode{}
	prev := &dum
	curr := head

	for curr.Next != nil {
		n := curr.Next
		curr.Next = prev
		prev = curr
		curr = n
	}

	curr.Next = prev
	head.Next = nil
	return curr
}
