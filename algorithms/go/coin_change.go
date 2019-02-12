package main

import (
	"fmt"
)

func main() {
	testCoins := []int{1, 2, 5}
	testAmount := 11

	result := coinChange(testCoins, testAmount)

	fmt.Println(result)
}

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	dp := make([]int, amount+1)

	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
		for _, c := range coins {
			if c > i {
				continue
			}

			if dp[i] > dp[i-c]+1 {
				dp[i] = dp[i-c] + 1
			}
		}
	}

	if dp[amount] > amount {
		return -1
	} else {
		return dp[amount]
	}
}
