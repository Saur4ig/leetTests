package stringsE

// 20. Valid Parentheses
// Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
// An input string is valid if:
//
// Open brackets must be closed by the same type of brackets.
// Open brackets must be closed in the correct order.
// Note that an empty string is also considered valid.
// link - https://leetcode.com/problems/valid-parentheses/
func isValid(s string) bool {
	a := &Stack{}
	for _, val := range s {
		switch string(val) {
		case "(":
			a.Push("(")
		case "[":
			a.Push("[")
		case "{":
			a.Push("{")
		case ")":
			if !isVal(a, "(") {
				return false
			}
		case "]":
			if !isVal(a, "[") {
				return false
			}
		case "}":
			if !isVal(a, "{") {
				return false
			}
		}
	}
	if a.IsEmpty() {
		return true
	}
	return false
}

type Stack struct {
	Value []string
	Size  int
}

func (s *Stack) Push(val string) {
	s.Value = append(s.Value, val)
	s.Size++
}

func (s *Stack) Pop() string {
	if s.Size == 0 {
		return ""
	}
	val := s.Value[s.Size-1]
	s.Value = s.Value[:s.Size-1]
	s.Size--
	return val
}

func (s *Stack) Top() string {
	if s.Size == 0 {
		return ""
	}
	return s.Value[s.Size-1]
}

func (s *Stack) IsEmpty() bool {
	return s.Size == 0
}

func isVal(st *Stack, pretender string) bool {
	if st.IsEmpty() {
		return false
	}
	r := st.Top()
	if r == pretender {
		st.Pop()
		return true
	}
	return false
}
