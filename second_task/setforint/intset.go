package intset

type Intsetstruct struct {
	elem map[int]bool
}

func (i Intsetstruct) Addtoset(num int) Intsetstruct {
	if i.elem == nil {
		var a = Intsetstruct{}
		a.elem = map[int]bool{}
		a.elem[num] = true
		return a
	}
	i.elem[num] = true
	return i
}
func (i Intsetstruct) Deletefromset(num int) Intsetstruct {
	if i.elem == nil {
		var a = Intsetstruct{}
		a.elem = map[int]bool{}
		return a
	}
	delete(i.elem, num)
	return i
}
func (i Intsetstruct) Checkforelem(num int) bool {
	if _, ok := i.elem[num]; ok {
		return true
	}
	return false
}
