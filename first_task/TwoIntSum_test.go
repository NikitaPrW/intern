package sum

import (
	"fmt"
	"log"
	"testing"
)

func TWoIntSumTest(t *testing.T) {
	log.SetFlags(log.Ltime | log.Lshortfile)
	var result []int
	testtargets := []int{1, 2, 3, 4, 5, 6, 7, 8}
	testarrays := [][]int{
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		[]int{1, 3, 5, 1, 2, 5, 6, 7, 234, 2, 6, 1, 4, 6, 7, 5},
	}
	var failsCount int

	for _, test := range testarrays {
		for index := 0; index < len(testtargets); index++ {
			result = TwoIntSum(test, testtargets[index])

			if len(result) != 2 {
				failsCount++
				log.Println("couldn't find numbers", test, testtargets[index])
			}
		}
	}
	if failsCount == len(testtargets) {
		fmt.Println("all test failed")
		t.Fail()
	}
}
