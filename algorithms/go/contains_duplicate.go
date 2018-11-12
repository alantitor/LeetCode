package main

import "fmt"

func main() {
	testCase := [...]int{1, 2, 3, 4}

	result := containsDuplicate(testCase[:])

	fmt.Println(result)
}

func containsDuplicate(nums []int) bool {
	myMap := make(map[int]bool, len(nums))

	for _, value := range nums {
		if myMap[value] {
			return true
		}
		myMap[value] = true
	}

	return false
}
