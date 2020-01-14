package main

import (
	"fmt"
)

//for a[i] range Slice

func twoIntSum(nums []int, target int) []int {
	var solution []int
	for i, value := range nums {
		for n := i + 1; n < len(nums); n++ {
			if value+nums[n] == target {
				solution = append(solution, i, n)
				return solution
			}
		}
	}
	fmt.Println(target)
}

func main() {
	a := []int{0, 1, 2, 3, 4}
	twoIntSum(a, 5)
}
