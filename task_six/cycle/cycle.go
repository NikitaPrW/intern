package cycle

type listNode struct {
	val  int
	next *listNode
}

func hasCycle(list *[]listNode) bool {
	pointerspassed := make(map[*listNode]bool)
	for _, value := range *list {
		_, ok := pointerspassed[value.next]
		if ok == true {
			return true
		} else {
			pointerspassed[value.next] = true
		}
	}
	return false
}
