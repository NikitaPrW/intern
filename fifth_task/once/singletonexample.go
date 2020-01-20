package singletone

import (
	"sync"
)

type mystruct struct {
	n int
}

var (
	executeonce sync.Once
)

func getInstance(a int) *mystruct {
	var newinstance mystruct
	executeonce.Do(func() {
		newinstance = mystruct{a}
	})
	return &newinstance
}
