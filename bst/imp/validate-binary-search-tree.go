package bst

import (
	"container/list"
	"fmt"
)

/*

98. Validate Binary Search Tree

Given the root of a binary tree, determine if it is a valid binary search tree (BST).

A valid BST is defined as follows:

The left
subtree
 of a node contains only nodes with keys less than the node's key.
The right subtree of a node contains only nodes with keys greater than the node's key.
Both the left and right subtrees must also be binary search trees.


Example 1:


Input: root = [2,1,3]
Output: true
Example 2:


Input: root = [5,1,4,null,null,3,6]
Output: false
Explanation: The root node's value is 5 but its right child's value is 4.


*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// O(N), O(logN) // iterative solution
func inOrderStack(root *TreeNode) bool {
	l := list.New()
	curr := root

	for curr != nil {
		l.PushBack(curr)
		curr = curr.Left
	}
	var prev *TreeNode
	for l.Len() > 0 {
		curr := l.Back()
		tmp := curr.Value.(*TreeNode)
		l.Remove(curr)
		fmt.Println(l, "-", tmp)
		if prev != nil && prev.Val >= tmp.Val {
			return false
		}
		// get the right child, and push the left child in stack
		ptr := tmp.Right
		for ptr != nil {
			l.PushBack(ptr)
			ptr = ptr.Left
		}
		prev = tmp
	}
	return true
}

// stack iterative ans recursive
// improve inorder traversal
func inOrder(root *TreeNode, ans *[]int) {
	if root == nil {
		return
	}

	inOrder(root.Left, ans)
	*ans = append(*ans, root.Val)
	inOrder(root.Right, ans)
}
func isValidBST1(root *TreeNode) bool {
	var ans []int
	inOrder(root, &ans)

	for i := 1; i < len(ans); i++ {
		if ans[i-1] >= ans[i] {
			return false
		}
	}

	return true
}
func isValidBST(root *TreeNode) bool {
	return inOrderStack(root)
}
