package linkedList

/*

19. Remove Nth Node From End of List

Given the head of a linked list, remove the nth node from the end of the list and return its head.

 

Example 1:


Input: head = [1,2,3,4,5], n = 2
Output: [1,2,3,5]
Example 2:

Input: head = [1], n = 1
Output: []
Example 3:

Input: head = [1,2], n = 1
Output: [1]
 

Constraints:

The number of nodes in the list is sz.
1 <= sz <= 30
0 <= Node.val <= 100
1 <= n <= sz
 

Follow up: Could you do this in one pass?
*/


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */


 func removeElement(h *ListNode, n int,count *int)*ListNode{
    if h == nil{
        return nil
    }
    next:= removeElement(h.Next,n,count)
    *count++
    if *count != n {
        h.Next = next
        return h
    }

    h.Next=nil
    return next
}
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
    var count int
    return removeElement(head,n,&count)
}