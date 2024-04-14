package binaryTree

/*

297. Serialize and Deserialize Binary Tree

Serialization is the process of converting a data structure or object into a sequence of bits so that it can be stored in a file or memory buffer, or transmitted across a network connection link to be reconstructed later in the same or another computer environment.

Design an algorithm to serialize and deserialize a binary tree. There is no restriction on how your serialization/deserialization algorithm should work. You just need to ensure that a binary tree can be serialized to a string and this string can be deserialized to the original tree structure.

Clarification: The input/output format is the same as how LeetCode serializes a binary tree. You do not necessarily need to follow this format, so please be creative and come up with different approaches yourself.



Example 1:


Input: root = [1,2,3,null,null,4,5]
Output: [1,2,3,null,null,4,5]
Example 2:

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

import (
	"fmt"
	"strconv"
	"strings"
)

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var ans []string
	serializeUtil(root, &ans)
	str := strings.Join(ans, ",")
	fmt.Println(str)
	return str

}

// it does a preorder traversal
func serializeUtil(root *TreeNode, ans *[]string) {
	if root == nil {
		*ans = append(*ans, "#")
		return
	}

	tmp := strconv.Itoa(root.Val)
	*ans = append(*ans, tmp)
	serializeUtil(root.Left, ans)
	serializeUtil(root.Right, ans)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	i := 0
	tmp := strings.Split(data, ",")
	root := deserializeUtil(tmp, &i)
	return root
}

// preorder deserializer
func deserializeUtil(data []string, i *int) *TreeNode {
	if data[*i] == "#" {
		return nil
	}

	tmp, _ := strconv.Atoi(data[*i])
	node := &TreeNode{Val: tmp}
	// increment the index to point to the left child
	*i++
	node.Left = deserializeUtil(data, i)
	// increment the index to point to the left child
	*i++
	node.Right = deserializeUtil(data, i)
	return node
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
