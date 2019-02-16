package main

import "fmt"

func main() {
	testCase := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

	result := maxSubArray(testCase)

	fmt.Println(result)
}

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums)+1)
	dp[0] = nums[0]
	maxValue := nums[0]

	for i := 1; i < len(nums); i++ {
		dp[i] = max(nums[i], dp[i-1]+nums[i])
		maxValue = max(maxValue, dp[i])
	}

	return maxValue
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
