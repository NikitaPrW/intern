package isomorphic

import (
	"fmt"
	"log"
	"testing"
)

func TestIsIsomorphic(t *testing.T) {
	log.SetFlags(log.Ltime | log.Lshortfile)
	teststrings := []string{
		"abc",
		"def",
		"egg",
		"add",
		"foo",
		"bar",
		"paper",
		"title",
	}
	var failsCount int

	for _, testS := range teststrings {
		for _, testT := range teststrings {
			result := isIsomorphic(testS, testT)
			smap := make(map[rune]int)
			tmap := make(map[rune]int)
			for _, value := range testS {
				smap[value] = smap[value] + 1 // создается map в котором хранится символ:количество вхождений в строке
			}
			for _, value := range testT {
				tmap[value] = tmap[value] + 1
			}
			var testresult = true
			numoflettersT := make(map[int]int)
			numoflettersS := make(map[int]int)
			for _, value := range smap {
				numoflettersS[value] = numoflettersS[value] + 1 //создается map, в котором хранятся количество вхождений в строке какого-либо символа:сколько таких символов, для которого такое количество вхождений
			}
			for _, value := range tmap {
				numoflettersT[value] = numoflettersT[value] + 1
			}
			if len(numoflettersS) > len(numoflettersT) {
				for iter := range numoflettersS {
					if numoflettersS[iter] != numoflettersT[iter] {
						testresult = false
					}
				}
			} else {
				for iter := range numoflettersT {
					if numoflettersS[iter] != numoflettersT[iter] {
						testresult = false
					}
				}
			}
			if testresult == result {
				log.Println("SUCCESS ", testT, testS, "result", "what it should be", result, "---", testresult)
				failsCount++
			}
			if testresult != result {
				log.Println("FAILED", testT, testS, "result", result, "testrelut", testresult)
			}
		}
	}
	if failsCount == len(teststrings) {
		fmt.Println("all test failed")
		t.Fail()
	}
}
