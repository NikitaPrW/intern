package singletone

import "testing"

import "fmt"

func TestSomefunction(t *testing.T) {
	instance1 := getInstance(1)
	fmt.Println(instance1.n)
	instance2 := getInstance(2)
	fmt.Println(instance2.n)
	if instance2.n == 0 {
		fmt.Println("ok")
	} else {
		fmt.Println("bad")
	}
}
