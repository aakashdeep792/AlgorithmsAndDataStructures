package binaryTree

/*
Given the roots of two binary trees root and subRoot, return true if there is a subtree of root with the same structure and node values of subRoot and false otherwise.

A subtree of a binary tree tree is a tree that consists of a node in tree and all of this node's descendants. The tree tree could also be considered as a subtree of itself.

Example 1:
Input: root = [3,4,5,1,2], subRoot = [4,1,2]
Output: true

Example 2:
Input: root = [3,4,5,1,2,null,null,null,null,0], subRoot = [4,1,2]
Output: false

*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSameTree1(root1, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}

	if root1 != nil && root2 != nil && root1.Val == root2.Val {
		return isSameTree(root1.Left, root2.Left) && isSameTree(root1.Right, root2.Right)
	}

	return false
}
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	tmp := isSameTree1(root, subRoot)
	if tmp {
		return tmp
	}

	if root == nil {
		return false
	}

	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)

}
