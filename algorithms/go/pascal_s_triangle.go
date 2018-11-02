package main

import "fmt"

func main() {
	testCase := 6

	ans := generate(testCase)

	fmt.Println(ans)
}

func generate(numRows int) [][]int {
	ret := make([][]int, numRows)

	if numRows == 0 {
		return ret
	}

	ret[0] = []int{1}

	for i := 1; i < numRows; i++ {
		ret[i] = make([]int, i+1)
		ret[i][0] = 1
		ret[i][i] = 1

		for j := 1; j <= len(ret[i])/2 && i > 1; j++ {
			ret[i][j] = ret[i-1][j-1] + ret[i-1][j]
			ret[i][len(ret[i])-j-1] = ret[i][j]
		}
	}

	return ret
}
