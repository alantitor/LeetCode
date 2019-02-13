package main

import "fmt"

func main() {
	testInput := []int{0, 1, 2, 0}

	result := minCostClimbingStairs(testInput)

	fmt.Println(result)
}

func minCostClimbingStairs(cost []int) int {
	dp := make([] int, len(cost)+1)

	for i := 2; i <= len(cost); i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}

	return dp[len(cost)]
}

func min(v1 int, v2 int) int {
	if v1 > v2 {
		return v2
	} else {
		return v1
	}
}
