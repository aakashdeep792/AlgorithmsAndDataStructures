package binaryTree

import "fmt"

/*

105. Construct Binary Tree from Preorder and Inorder Traversal

Given two integer arrays preorder and inorder where preorder is the preorder traversal of a binary tree and inorder is the inorder traversal of the same tree, construct and return the binary tree.



Example 1:


Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
Output: [3,9,20,null,null,15,7]
Example 2:

Input: preorder = [-1], inorder = [-1]
Output: [-1]


*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func createNode(val int, left, right *TreeNode) *TreeNode {
	return &TreeNode{
		Val:   val,
		Left:  left,
		Right: right,
	}
}

type TreeBuilder struct {
	po        []int
	io        []int
	tmap      map[int]int
	rootIndex *int
}

func (t *TreeBuilder) buildTreeUtil(l, r int) *TreeNode {
	// all the incorrect range scenario
	if l > r || l < 0 || r >= len(t.io) {
		return nil
	}

	// in case when l==r, update the rootIndex
	rootIndex := *(t.rootIndex)
	// if rootIndex < 0 || rootIndex >= len(t.io) {
	// 	return nil
	// }

	*(t.rootIndex)++

	// leaf node
	if l == r {
		return createNode(t.io[l], nil, nil)
	}
	// internal node
	m := t.tmap[t.po[rootIndex]]
	left := t.buildTreeUtil(l, m-1)
	right := t.buildTreeUtil(m+1, r)

	return createNode(t.po[rootIndex], left, right)
}

func BuildTree(preorder []int, inorder []int) *TreeNode {
	tmap := make(map[int]int)
	for i, v := range inorder {
		tmap[v] = i
	}
	fmt.Println(tmap)
	rootIndex := 0
	t := &TreeBuilder{
		po:        preorder,
		io:        inorder,
		tmap:      tmap,
		rootIndex: &rootIndex,
	}
	return t.buildTreeUtil(0, len(preorder)-1)
}

/*
Implement this solution sometimes

The basic idea is here:
Say we have 2 arrays, PRE and IN.
Preorder traversing implies that PRE[0] is the root node.
Then we can find this PRE[0] in IN, say it's IN[5].
Now we know that IN[5] is root, so we know that IN[0] - IN[4] is on the left side, IN[6] to the end is on the right side.
Recursively doing this on subarrays, we can build a tree out of it :)

*/
