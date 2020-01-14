package sum

import (
	"fmt"
)

func TwoIntSum(nums []int, target int) []int {
	var solution []int
	for i, value := range nums {
		for n := i + 1; n < len(nums); n++ {
			if value+nums[n] == target {
				solution = append(solution, i, n)
				return solution
			}
		}
	}
	fmt.Println("couldn't find two numbers")
	return solution
}
