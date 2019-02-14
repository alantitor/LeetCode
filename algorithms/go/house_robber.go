package main

import "fmt"

func main() {
	testCase := []int{2, 7, 9, 3, 1}

	result := rob(testCase)

	fmt.Println(result)
}

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(dp[0], nums[1])

	for i := 2; i < len(nums); i ++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}

	return dp[len(nums)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
