/*
	Given the roots of two binary trees p and q, write a function to check if they are the same or not.
	Two binary trees are considered the same if they are structurally identical, and the nodes have the same value.

	Example 1:
	Input: p = [1,2,3], q = [1,2,3]
	Output: true

	Example 2:
	Input: p = [1,2], q = [1,null,2]
	Output: false

	Example 3:
	Input: p = [1,2,1], q = [1,1,2]
	Output: false

	Constraints:
		The number of nodes in both trees is in the range [0, 100].
		-104 <= Node.val <= 104

	if I can traverse both at the same time, can check structure and values during the traverse
	several checks can be performed at each node, and possible return false,
	otherwise, we return the result of checking left
	then of checking right
	finally, just return true if we've made it all the way to the end
*/

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p == nil || q == nil || p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
