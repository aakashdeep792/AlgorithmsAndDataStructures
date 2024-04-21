package linkedList

/*
206. Reverse Linked List
Solved
Easy
Topics
Companies
Given the head of a singly linked list, reverse the list, and return the reversed list.



Example 1:


Input: head = [1,2,3,4,5]
Output: [5,4,3,2,1]
Example 2:


Input: head = [1,2]
Output: [2,1]
Example 3:

Input: head = []
Output: []

*/

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseListItr(head *ListNode) *ListNode {
	curr := head
	var newH *ListNode
	for curr != nil {
		tmp := curr
		curr = curr.Next
		tmp.Next = newH
		newH = tmp

	}
	return newH
}

func reverseListRecursive(head *ListNode) (*ListNode, *ListNode) {
	if head == nil {
		return nil, nil
	}
	if head.Next == nil {
		return head, head
	}

	h, t := reverseListRecursive(head.Next)

	t.Next = head // the head become the new tail
	head.Next = nil
	return h, head
}

func reverseList(head *ListNode) *ListNode {
	h := reverseListItr(head)
	return h
}
