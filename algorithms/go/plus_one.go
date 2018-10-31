package main

import "fmt"

func main() {
	testCase := [...]int{9, 8, 9}

	ans := plusOne(testCase[:])

	fmt.Println(ans)
}

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			break
		} else {
			digits[i] = 0
		}
	}

	if digits[0] == 0 {
		digits = append([]int{1}, digits...)
	}

	return digits
}
