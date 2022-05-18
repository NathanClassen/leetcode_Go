/*
	You are given the heads of two sorted linked lists list1 and list2.
	Merge the two lists in a one sorted list. The list should be made by splicing together the nodes of the
		first two lists.

	Return the head of the merged linked list.



	Example 1:
	Input: list1 = [1,2,4], list2 = [1,3,4]
	Output: [1,1,2,3,4,4]

	Example 2:
	Input: list1 = [], list2 = []
	Output: []

	Example 3:
	Input: list1 = [], list2 = [0]
	Output: [0]

	Constraints:
		The number of nodes in both lists is in the range [0, 50].
		-100 <= Node.val <= 100
		Both list1 and list2 are sorted in non-decreasing order.
*/

/*
		given two lists
			create one new list, made of both given lists with nodes sorted

		inputs are sortet

		inputs: two lists
		outputs: one list

		list1 = [1]

		list2 = [1,3,4]

		mental model:

		traverse both lists, adding the smaller of the two current nodes to the end of the new list
		if either list runs out of nodes, safe to attach remaining nodes of other list

		two empty lists? -- return nil
		one empty list? -- return other list
		         .
		1  2  4  nil
		      .
		1  3  5  nil

	               e
	d  1  1  2  3  4  5  nil

		create dummy node (will eventually return dummy.Next)
		end = *dummy
		handle edge cases
		set current nodes
		loop
			if either current node is nil
				break
			if lNode <= rNode
				newListEnd.Next = lNode
				newListEnd = lNode
				currL = currL.Next
			else
				end.Next = rNode
				end = rNode
				currR = currR.Next

		if currR == nil
			newListEnd.Next = currL
		else if currL == nil
			newListEnd.Next = currR

		return dummy.Next

*/

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := ListNode{}
	end := &dummy
	for {
		if list1 == nil || list2 == nil {
			break
		}
		if list1.Val <= list2.Val {
			end.Next = list1
			end = list1
			list1 = list1.Next
		} else {
			end.Next = list2
			end = list2
			list2 = list2.Next
		}
	}

	if list1 == nil {
		end.Next = list2
	} else if list2 == nil {
		end.Next = list1
	}

	return dummy.Next
}
