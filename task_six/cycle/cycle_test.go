package cycle

import (
	"fmt"
	"testing"
)

func TestHascycle(t *testing.T) {
	//head := []int{2, 4, 6, 8, 1}
	//pos := 1
	head := []int{2}
	pos := -1
	list := make([]listNode, len(head))
	i := len(head) - 1
	if pos >= 0 && pos < len(head) {
		list[i] = listNode{val: head[i], next: &list[pos]}
	}
	for i = 0; i < len(list)-1; i++ {
		list[i] = listNode{val: head[i], next: &list[i+1]}

	}
	iscycle := hasCycle(&list)
	fmt.Println(iscycle)

}
