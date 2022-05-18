/*
	Given the root of a binary tree, return its maximum depth.

	A binary tree's maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.

	Example 1:
	Input: root = [3,9,20,null,null,15,7]
	Output: 3

	Example 2:
	Input: root = [1,null,2]
	Output: 2

	Constraints:
		The number of nodes in the tree is in the range [0, 104].
		-100 <= Node.val <= 100
*/
/*
	how to travel tree,
		same as linked list, but with left, and right instead of next
		looping would be very complicated as I would have to track all possible routes manually,
			obviously must use recursion

	base case is a null node, return 0
	otherwise you have a node that adds to total
		then you must count left nodes from that location

	base case, single node, left and right is null, return 1, check after recursive call i think
	in case of node with left and no right, would say
		if single node return 1 (or add to toal, whatevr... need to figure out the combine step
	if left, (this will be depth first) call fun passing in node to left and current total
	if right, call func passing in node to right and total

	actually, probably add 1 first, then pass to next left node if possible
	otherwise set/compare new left total,
	then do right

	just need to have a current length tracker and a global
	whether right or left search, curr gets bumped until the return, when we get to a cross road, compare/set curr/global
	so if there is even one head node, curr should be set to 1






*/

package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return int(math.Max(float64(1+maxDepth(root.Left)), float64(1+maxDepth(root.Right))))
}
