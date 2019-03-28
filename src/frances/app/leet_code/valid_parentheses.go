/*
Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Note that an empty string is also considered valid.

Example 1:

Input: "()"
Output: true
Example 2:

Input: "()[]{}"
Output: true
Example 3:

Input: "(]"
Output: false
*/
package main

type stack struct {
	el []rune
}

func (st *stack) push(b rune) {
	st.el = append(st.el, b)
}

//先进后出
func (st *stack) pop() (rune, bool) {
	if len(st.el) > 0 {
		popEl := st.el[len(st.el)-1]
		st.el = st.el[:len(st.el)-1]
		return popEl, true
	}
	return 0, false
}

func validParentheses(str string) bool {
	var parenthesesMap = map[rune]rune {
		')':'(',
		'}':'{',
		']':'[',
	}

	var s stack
	for _, v := range str {
		switch v {
		case '(','{','[':
			s.push(v)
		case ')', '}', ']':
			if pop, ok := s.pop();!ok || pop != parenthesesMap[v] {
				return false
			}
		}
	}

	if len(s.el) > 0 {
		return false
	}

	return true
}
func main() {
	println(validParentheses("({})"))
}
