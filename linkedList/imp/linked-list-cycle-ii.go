package linkedList

/*

142. Linked List Cycle II

Given the head of a linked list, return the node where the cycle begins. If there is no cycle, return null.

There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the next pointer. Internally, pos is used to denote the index of the node that tail's next pointer is connected to (0-indexed). It is -1 if there is no cycle. Note that pos is not passed as a parameter.

Do not modify the linked list.



Example 1:


Input: head = [3,2,0,-4], pos = 1
Output: tail connects to node index 1
Explanation: There is a cycle in the linked list, where tail connects to the second node.
Example 2:


Input: head = [1,2], pos = 0
Output: tail connects to node index 0
Explanation: There is a cycle in the linked list, where tail connects to the first node.
Example 3:


Input: head = [1], pos = -1
Output: no cycle
Explanation: There is no cycle in the linked list.

*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func detectCycle(head *ListNode) *ListNode {
	// base case
	if head == nil || head.Next == nil {
		return nil
	}

	ptr1, ptr2 := head, head
	f := true
	//check ptr1 != nil can be skpped as ptr2 will reach the end first
	for ptr2 != nil && ptr2.Next != nil {
		if ptr1 == ptr2 && f == false {
			break
		}

		f = false
		ptr1 = ptr1.Next
		ptr2 = ptr2.Next.Next
	}

	if ptr1 != ptr2 {
		return nil
	}
	// we know that cycle exits, now change ptr1 to point to head of the list
	ptr1 = head
	for ptr1 != ptr2 {
		ptr1 = ptr1.Next
		ptr2 = ptr2.Next
	}

	return ptr1
}
