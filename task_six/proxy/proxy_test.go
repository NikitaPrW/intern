package proxy

import (
	"fmt"
	"testing"
)

func TestProxyimplementation(t *testing.T) {
	intern := new(intern)
	salary := intern.CalculateSalary()
	fmt.Println("intern will have", salary)
	programmer := new(programmer)
	salary = programmer.CalculateSalary()
	fmt.Println("developer will have", salary)
}
