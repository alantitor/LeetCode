package main

import "fmt"

func main() {
	testCase := 3

	ans := getRow(testCase)

	fmt.Println(ans)
}

func getRow(rowIndex int) []int {
	ret := make([]int, rowIndex+1)
	ret[0] = 1

	if rowIndex == 0 {
		return ret
	}

	ret[rowIndex] = 1

	for i := 1; i < rowIndex+1; i ++ {
		if i == 1 {
			ret[1] = 1
		}
		mid := i / 2
		for j := mid; j > 0; j-- {
			ret[j] += ret[j-1]
			ret[2*mid-j+i%2] = ret[j]
		}
	}

	return ret
}
