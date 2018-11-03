package main

import (
	"fmt"
)

func main() {
	testCaseA := "1011"
	testCaseB := "1011"

	ans := addBinary(testCaseA, testCaseB)

	fmt.Println(ans)
}

func addBinary(a string, b string) string {
	ret := ""
	indexA := len(a) - 1
	indexB := len(b) - 1
	carry := byte(0)

	for indexA >= 0 || indexB >= 0 || carry == 1 {
		if indexA >= 0 {
			carry += a[indexA] - '0'
		}
		if indexB >= 0 {
			carry += b[indexB] - '0'
		}

		ret = string(carry%2+'0') + ret

		indexA--
		indexB--
		carry /= 2
	}

	return ret
}
