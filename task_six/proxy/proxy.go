package proxy

import (
	"fmt"
)

type developer interface {
	CalculateSalary() int
}
type intern struct { //суррогат
	programmer developer
}

func (i *intern) CalculateSalary() int {
	if i.programmer == nil {
		i.programmer = &programmer{}
	}
	fmt.Println("I'm an intern")
	return 50
}

type programmer struct { //основной
}

func (p *programmer) CalculateSalary() int {
	fmt.Println("I'm the real dveloper")
	return 100
}
