package main

import (
	"fmt"
)

func main() {
	testCase := ""

	result := firstUniqChar(testCase)

	fmt.Println(result)
}

func firstUniqChar(s string) int {
	countMap := make(map[int]int, 26)

	for _, value := range s {
		countMap[int(value)] += 1
	}

	for index, value := range s {
		if countMap[int(value)] == 1 {
			return index
		}
	}

	return -1
}
