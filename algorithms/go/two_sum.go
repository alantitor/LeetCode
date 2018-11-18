package main

import "fmt"

func main() {
	testNumbers := []int{1, 2, 9, 14}
	testTarget := 16

	result := twoSum(testNumbers, testTarget)

	fmt.Println(result)
}

func twoSum(nums []int, target int) []int {
	result := [2]int{}
	numMap := make(map[int]int, len(nums))

	for index, value := range nums {
		_, isExist := numMap[target-value]

		if isExist {
			result[0] = numMap[target-value]
			result[1] = index
			break
		} else {
			numMap[value] = index
		}
	}

	return result[:]
}
