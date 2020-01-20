package bracketchecker

// функция проверяет корректность открытых/закрытых скобок
func isValid(s *string) bool {

	bracketlist := map[rune]rune{
		'[': ']',
		'{': '}',
		'(': ')',
	}
	stack := make([]rune, len(*s))

	esi := 0
	for _, bracket := range *s {

		if bracket == '{' || bracket == '[' || bracket == '(' { //если встречается открывающая скобка, то она записывается на стэк, и указатель стэка увеличивается на 1
			stack[esi] = bracketlist[bracket]
			esi++
		}

		if bracket == '}' || bracket == ']' || bracket == ')' { // проверка:если найдена закрывающая скобка, то проверка на то, что предыдущая скобка была того же типа и открывающая, если все ок, то указатель на стэк ументшается
			if stack[esi-1] == bracket && esi > 0 {
				esi--
			} else {
				return false
			}

		}

	}
	return true
}
