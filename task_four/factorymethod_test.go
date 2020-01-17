package main

import (
	"fmt"
	"testing"

	"github.com/NikitaPrW/intern/task_four/factorymethod"
)

func TestFactory(t *testing.T) {
	wantrateresults := []int{1, 2}
	factory := factorymethod.NewCreator()
	fmt.Println(factory)
	product := []factorymethod.Producter{
		factory.CreateProduct("pc"),
		factory.CreateProduct("monitor"),
	}
	rateresults := make([]int, 2)
	for i, v := range product {
		rateresults[i] = v.Rate()
		fmt.Println(rateresults[i])
	}
	for i := 0; i < len(rateresults); i++ {
		if rateresults[i] != wantrateresults[i] {
			t.Errorf("false rate")
		}
	}

}
