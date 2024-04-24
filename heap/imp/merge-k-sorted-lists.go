package heap

import "container/heap"

/*
23. Merge k Sorted Lists

You are given an array of k linked-lists lists, each linked-list is sorted in ascending order.

Merge all the linked-lists into one sorted linked-list and return it.



Example 1:

Input: lists = [[1,4,5],[1,3,4],[2,6]]
Output: [1,1,2,3,4,4,5,6]
Explanation: The linked-lists are:
[
  1->4->5,
  1->3->4,
  2->6
]
merging them into one sorted list:
1->1->2->3->4->4->5->6
Example 2:

Input: lists = []
Output: []
Example 3:

Input: lists = [[]]
Output: []

*/

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

type listObj struct {
	count int
	head  *ListNode
}

func merge(l, m *listObj) *listObj {
	curr1, curr2 := l.head, m.head
	var head, ptr *ListNode
	for curr1 != nil && curr2 != nil {
		var tmp *ListNode
		if curr1.Val > curr2.Val {
			tmp = curr2
			curr2 = curr2.Next
		} else {
			tmp = curr1
			curr1 = curr1.Next
		}

		tmp.Next = nil

		if head == nil {
			head = tmp
			ptr = head
		} else {
			ptr.Next = tmp
			ptr = ptr.Next
		}
	}
	// when one/both input list is empty
	if head == nil {
		if curr1 == nil {
			head = curr2
		} else {
			head = curr1
		}
	} else { // after mergeing one/both list become empty
		if curr1 == nil {
			ptr.Next = curr2
		} else {
			ptr.Next = curr1
		}
	}

	return &listObj{count: l.count + m.count, head: head}
}

type priorityQueue []*listObj

func (p *priorityQueue) Len() int {
	return len(*p)
}

func (p *priorityQueue) Less(i, j int) bool {
	//min heap
	// i<j && *p[i].count<*p[j].count
	return (*p)[i].count < (*p)[j].count // true mean no swap
}

func (p *priorityQueue) Swap(i, j int) {
	(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
}

func (p *priorityQueue) Push(val any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	tmp := val.(*listObj)
	*p = append(*p, tmp)
}

func (p *priorityQueue) Pop() any {
	n := len(*p)
	val := (*p)[n-1]
	*p = (*p)[0 : n-1]
	return val
}

func AddCount(head *ListNode) *listObj {
	curr := head
	count := 0
	for curr != nil {
		curr = curr.Next
		count++
	}

	return &listObj{count: count, head: head}
}

func MergeKLists(lists []*ListNode) *ListNode {
	var p priorityQueue
	for _, l := range lists {
		p = append(p, AddCount(l))
	}

	// Initialise Heap
	heap.Init(&p)

	for p.Len() > 0 {
		var x, y *listObj
		x = heap.Pop(&p).(*listObj)
		// if heap become empty return the list
		if p.Len() > 0 {
			y = heap.Pop(&p).(*listObj)
		} else {
			return x.head
		}

		tmp := merge(x, y)
		heap.Push(&p, tmp)
	}

	return nil
}
