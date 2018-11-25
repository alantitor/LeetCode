package main

import "fmt"

func main() {
	testCase := []int{2, -1, 1, 2, 2}

	result := circularArrayLoop(testCase)

	fmt.Println(result)
}

func circularArrayLoop(nums []int) bool {
	if len(nums) == 0 {
		return false
	}

	result := false
	length := len(nums)
	count := 0
	index := 0
	direction := nums[0]

	for count < length {
		newIndex := (index + nums[index] + length) % length

		if direction*nums[newIndex] < 0 {
			result = false
			break
		}

		if newIndex == 0 {
			result = true
			break
		}

		index = newIndex
		count++
	}

	return result
}
