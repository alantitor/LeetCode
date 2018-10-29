package main

import "fmt"

func main() {
	testCase := [...]int{-1, -1, 0, 1, 1, 0}
	//testCase := [...]int{1, 7, 3, 6, 5, 6}

	ans := pivotIndex(testCase[:])

	fmt.Println("ans: ", ans) // 0
}

func pivotIndex(nums []int) int {
	pi, sum, leftSum := -1, 0, 0

	for _, v := range nums {
		sum += v
	}

	for i := 0; i < len(nums); i++ {
		if leftSum == sum-nums[i]-leftSum {
			pi = i
			break
		}
		leftSum += nums[i]
	}

	return pi
}
