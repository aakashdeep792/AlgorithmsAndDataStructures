package strings

import "container/list"

/*

20. Valid Parentheses

Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Every close bracket has a corresponding open bracket of the same type.


Example 1:

Input: s = "()"
Output: true
Example 2:

Input: s = "()[]{}"
Output: true
Example 3:

Input: s = "(]"
Output: false

*/

// using stack
func IsValid(s string) bool {
	n := len(s)
	l := list.New()

	for i := 0; i < n; i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			l.PushFront(s[i])
		} else {
			if l.Len() > 0 {
				tmp := l.Front()
				top := tmp.Value.(byte)

				if (s[i] == ')' && top == '(') || (s[i] == '}' && top == '{') ||
					(s[i] == ']' && top == '[') {
					l.Remove(tmp)
				} else {
					return false
				}
			} else {
				return false
			}

		}
	}

	if l.Len() > 0 {
		return false
	}
	return true
}
