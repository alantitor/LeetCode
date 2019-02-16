package main

import "fmt"

func main() {
	testCase := []int{7, 2, 5, 1, 6, 4}

	result := maxProfit(testCase)

	fmt.Println(result)
}

func maxProfit(prices []int) int {
	currMax := 0
	totalMax := 0

	for i := 1; i < len(prices); i++ {
		currMax = max(0, currMax+prices[i]-prices[i-1])
		totalMax = max(totalMax, currMax)
	}

	return totalMax
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
