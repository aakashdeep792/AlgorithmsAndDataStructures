package bst

/*
230. Kth Smallest Element in a BST

Given the root of a binary search tree, and an integer k, return the kth smallest value (1-indexed) of all the values of the nodes in the tree.



Example 1:


Input: root = [3,1,4,null,2], k = 1
Output: 1
Example 2:


Input: root = [5,3,6,2,4,null,null,1], k = 3
Output: 3

*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// Similar to this approach can be applied to the valid Bst by keeping track of prev variable
func inOrderSmallest(root *TreeNode, k int, count *int) *TreeNode {
	if root == nil {
		return nil
	}

	left := inOrderSmallest(root.Left, k, count)
	(*count)++
	if *count == k {
		return root
	}

	right := inOrderSmallest(root.Right, k, count)
	if left == nil {
		return right
	}

	return left

}

func kthSmallest(root *TreeNode, k int) int {
	count := 0
	tmp := inOrderSmallest(root, k, &count)
	if tmp == nil {
		return -1
	}
	return tmp.Val

}
