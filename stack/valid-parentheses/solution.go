package validparentheses

import "errors"

func isValid(s string) bool {
	stk := Stack{}
	for _, sk := range s {
		skstr := string(sk)
		
		if skstr == "(" || skstr == "[" || skstr == "{" {
			stk.push(skstr)
			continue
		}

		last := stk.last()
		
		switch {
		case last == "(":
			if skstr != ")" {
				return false
			}

			break

		case last == "{":
			if skstr != "}" {
				return false
			}

			break

		case last == "[":
			if skstr != "]" {
				return false
			}

			break
		}

	}

	if len(stk.stack) !=0 {
		return false
	}
	return true
}

type Stack struct {
	stack []string
}

func (s *Stack) last()( string, error) {
	n := len(s.stack)

	if n == 0 {
		return "", errors.New("last of empty stack")
	}
	l := s.stack[n-1]
	s.stack = s.stack[:n-1]
	return l, nil
}

func (s *Stack) push(sk string) {
	s.stack = append(s.stack, sk)
}
