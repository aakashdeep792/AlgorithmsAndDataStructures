package bst

/*

99 Recover binary search tree

You are given the root of a binary search tree (BST), where the values of exactly two nodes of the tree were swapped
 by mistake. Recover the tree without changing its structure.

Input: root = [1,3,null,null,2]
Output: [3,1,null,null,2]
Explanation: 3 cannot be a left child of 1 because 3 > 1. Swapping 1 and 3 makes the BST valid.


Input: root = [3,1,4,null,null,2]
Output: [2,1,4,null,null,3]
Explanation: 2 cannot be in the right subtree of 3 because 2 < 3. Swapping 2 and 3 makes the BST valid.

*/

// * Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func IdentifySwappedNodes(root *TreeNode, first, sec, prev **TreeNode) {
	if root == nil {
		return
	}

	IdentifySwappedNodes(root.Left, first, sec, prev)

	if (*prev) != nil && (*prev).Val > root.Val {
		if *first == nil {
			*first = *prev
		}
		*sec = root // this handle the adjecent case
	}

	*prev = root

	IdentifySwappedNodes(root.Right, first, sec, prev)

}

func RecoverTree(root *TreeNode) {
	var first, sec, prev *TreeNode
	IdentifySwappedNodes(root, &first, &sec, &prev)
	(*first).Val, (*sec).Val = (*sec).Val, (*first).Val
}
