package bracketchecker

func isValid(s *string) bool {

	bracketlist := map[rune]rune{
		'[': ']',
		'{': '}',
		'(': ')',
	}
	stack := make([]rune, len(*s))

	esi := 0
	for _, bracket := range *s {

		if bracket == '{' || bracket == '[' || bracket == '(' {
			stack[esi] = bracketlist[bracket]
			esi++
		}

		if bracket == '}' || bracket == ']' || bracket == ')' {
			if stack[esi-1] == bracket && esi > 0 {
				esi--
			} else {
				return false
			}

		}

	}
	return true
}
