package intset

type intsetstruct struct {
	elem map[int]bool
}

func (i *intsetstruct) Addtoset(num int) intsetstruct {
	if i.elem == nil {
		var a = intsetstruct{}
		a.elem = map[int]bool{}
		a.elem[num] = true
		return a
	}
	i.elem[num] = true
	return *i
}
func (i *intsetstruct) Deletefromset(num int) intsetstruct {
	if i.elem == nil {
		var a = intsetstruct{}
		a.elem = map[int]bool{}
		return a
	}
	delete(i.elem, num)
	return *i
}
func (i *intsetstruct) Checkforelem(num int) bool {
	if _, ok := i.elem[num]; ok {
		return true
	}
	return false
}
