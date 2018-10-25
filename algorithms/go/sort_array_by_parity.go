package main

import (
	"fmt"
)

func main() {
	testCase := [...]int{3, 1, 2, 4}

	ans := sortArrayByParity(testCase[:])

	fmt.Println(ans)
}

func sortArrayByParity(A []int) []int {
	ret := make([]int, len(A))
	evenIndex := 0
	oddIndex := len(A) - 1

	for _, v := range A {
		if isEven(v) {
			ret[evenIndex] = v
			evenIndex++
		} else {
			ret[oddIndex] = v
			oddIndex--
		}
	}

	return ret
}

func isEven(num int) bool {
	return num%2 == 0
}
