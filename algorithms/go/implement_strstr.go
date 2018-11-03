package main

import "fmt"

func main() {
	haystack := ""
	needle := ""

	ans := strStr(haystack, needle)

	fmt.Println(ans)
}

func strStr(haystack string, needle string) int {
	firstIndex := -1
	findFlag := false

	if len(needle) == 0 {
		return 0
	}

	for i := 0; i <= len(haystack)-len(needle); i++ {
		for j := 0; j < len(needle); j++ {
			if haystack[i+j] == needle[j] {
				findFlag = true
			} else {
				findFlag = false
				break
			}
		}

		if findFlag {
			firstIndex = i
			break
		}
	}

	return firstIndex
}
