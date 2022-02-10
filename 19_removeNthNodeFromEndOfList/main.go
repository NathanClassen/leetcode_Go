package main

import (
	"fmt"
)

/*
	Given the head of a linked list, remove the nth node from the end of the list and return its head.

	Example 1:
	Input: head = [1,2,3,4,5], n = 1
	Output: [1,2,3,5]

	Example 2:
	Input: head = [1], n = 1
	Output: []

	Example 3:
	Input: head = [1,2], n = 1
	Output: [1]

	Definition for singly-linked list.
	type ListNode struct {
	    Val int
	    Next *ListNode
	}

	Constraints:
		The number of nodes in the list is sz.
		1 <= sz <= 30
		0 <= Node.val <= 100
		1 <= n <= sz



	need to find the:
		- nth node
		- nth-1 node
		- nth+1 node

	iterate list, tracking each node in a slice of *node
	once length is found, we know what the above nodes will be and can fetch in constant time lookup from slice
	then connect nth-1 node to nth+1 node
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := ListNode{}
	curr := &head
	for _, val := range []int{1, 2, 3} {
		newNode := ListNode{
			Val:  val,
			Next: nil,
		}
		curr.Next = &newNode
		curr = &newNode
	}
	head = *removeNthFromEnd(head.Next, 3)

	curr = &head
	for {
		fmt.Printf("%d  ", curr.Val)
		curr = curr.Next
		if curr == nil {
			fmt.Println()
			break
		}
	}
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var nodes []*ListNode
	currNode := head

	for {
		nodes = append(nodes, currNode)
		currNode = currNode.Next
		if currNode == nil {
			break
		}
	}

	if len(nodes) == 1 {
		return &ListNode{}
	}

	preIdx, postIdx := len(nodes)-(n+1), len(nodes)-(n-1)
	var preNode, postNode *ListNode

	if preIdx < 0 {
		return nodes[1]
	}
	preNode = nodes[len(nodes)-(n+1)]

	if postIdx >= len(nodes) {
		postNode = nil
	} else {
		postNode = nodes[len(nodes)-(n-1)]
	}

	preNode.Next = postNode
	return head
}
