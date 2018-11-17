package main

import (
	"fmt"
)

func main() {

	testCase := 19

	result := isHappy(testCase)

	fmt.Println(result)
}

func isHappy(n int) bool {
	result := true
	record := make(map[int]bool)
	currValue := n

	for currValue != 1 {
		sum := 0

		for ; currValue > 0; currValue /= 10 {
			tempValue := currValue % 10
			sum += tempValue * tempValue
		}

		if record[sum] {
			result = false
			break
		} else {
			record[sum] = true
			currValue = sum
		}
	}

	return result
}
