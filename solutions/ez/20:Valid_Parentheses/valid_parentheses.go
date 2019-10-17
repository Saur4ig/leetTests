package _0_Valid_Parentheses

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
	return a.IsEmpty()
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
