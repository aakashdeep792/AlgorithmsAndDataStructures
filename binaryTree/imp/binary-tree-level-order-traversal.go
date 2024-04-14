package binaryTree

import "container/list"

/*
102. Binary Tree Level Order Traversal

Given the root of a binary tree, return the level order traversal of its nodes' values. (i.e., from left to right, level by level).



Example 1:


Input: root = [3,9,20,null,null,15,7]
Output: [[3],[9,20],[15,7]]
Example 2:

Input: root = [1]
Output: [[1]]
Example 3:

Input: root = []
Output: []

*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func LevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	q := list.New()
	q.PushBack(root)
	var ans [][]int
	for q.Len() > 0 {
		qLen := q.Len()
		var tmp []int
		for i := 0; i < qLen; i++ {
			ele := q.Front()
			node := ele.Value.(*TreeNode)
			tmp = append(tmp, node.Val)
			if node.Left != nil {
				q.PushBack(node.Left)
			}

			if node.Right != nil {
				q.PushBack(node.Right)
			}
			q.Remove(ele)
		}
		ans = append(ans, tmp)
	}

	return ans
}
