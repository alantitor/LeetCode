package main

import "fmt"

func main() {
	j := "aAb"
	s := "aAAbbbb"

	ans := numJewelsInStones(j, s)

	fmt.Println(ans)
}

func numJewelsInStones(J string, S string) int {
	ret := 0
	jewel := make(map[string]bool)

	for _, elem := range J {
		jewel[string(elem)] = true
	}

	for _, elem := range S {
		if jewel[string(elem)] {
			ret += 1
		}
	}

	return ret
}
