package main

import (
	"fmt"
)

func main() {
	testCase1 := []int{1, 2, 2, 1}
	testCase2 := []int{2, 2}

	ret := intersection(testCase1[:], testCase2[:])

	fmt.Println(ret)
}

func intersection(nums1 []int, nums2 []int) []int {
	record := make(map[int]bool)
	var result []int

	for _, value := range nums1 {
		record[value] = true
	}

	for _, value := range nums2 {
		if record[value] {
			result = append(result, value)
			record[value] = false
		}
	}

	return result
}
