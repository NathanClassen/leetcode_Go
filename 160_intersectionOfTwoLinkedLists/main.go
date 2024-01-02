/*
	Given the heads of two singly linked-lists headA and headB, return the node at which the two lists intersect.
	If the two linked lists have no intersection at all, return null.

	For example, the following two linked lists begin to intersect at node c1:
                      
	 A  
	[a1]-->[a2]-->[c1]-->[c2]-->[c3]-->nil
                   ^
	               ^
	              [b1]
                   B

	brute force:

	two pointers, one at aHead one at bHead
		vars of type node
	advance aHead fast (inner), advance bHead slow (outer)
		outer loops as long as bHead.val is NOT nil. bHead is moved along at end of outer loop
		inner loops as long as aHead.val is NOT nil, moved along at end of loop
	if aHead.val == bHead.next then that is the intersection, return aHead.val (or bHead.next)
	if bHead pointer gets to it's nil val, then there is no interseciton, return nil

	Follow up: Could you write a solution that runs in O(m + n) time and use only O(1) memory?
	
*/

package main

type ListNode struct {
	Val int
	Next *ListNode
}


func getIntersectionNode(headA, headB *ListNode) *ListNode {
	aHead := headA
    for {
		for {
			if headA == headB {
				return headA
			}

			if headA = headA.Next; headA == nil {
				break
			}
		}
		if headB = headB.Next; headB == nil {
			break
		}
		headA = aHead
	}
	return nil
}