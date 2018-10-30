package main

import (
	"fmt"
)

func main() {
	testCase := [...]int{0, 0, 3, 2}

	ans := dominantIndex(testCase[:])

	fmt.Println("ans: ", ans)
}

func dominantIndex(nums []int) int {
	var first, second, index int

	for i, v := range nums {
		if v > first {
			index = i
			second = first
			first = v
		} else if v > second {
			second = v
		}
	}

	if first < second * 2 {
		index = -1
	}

	return index
}
