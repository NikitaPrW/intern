package isomorphic

// isIsomorphic принимает на вход две строки и возвращает ответ, являются ли они изоморфными
func isIsomorphic(s string, t string) bool {
	answer := false
	if len(s) == len(t) { //проврека на равную длинну, так как в задании указано, что строки равной длинны
		smap := make(map[rune]int)
		tmap := make(map[rune]int)
		for _, value := range s { //записываются все буквы в map для строки s
			smap[value] = 0
		}
		for _, value := range t { //записываются все буквы в map для строки t
			tmap[value] = 0
		}
		if len(smap) == len(tmap) { // если длина двух map'ов совпадает, то строки изоморфны
			answer = true
		}
		return answer
	}
	return answer
}
