package main

import "fmt"

func main() {
	testNums1 := []int{1, 2, 3}
	testNums2 := []int{2, 3, 3, 3, 4}

	result := intersect(testNums1, testNums2)

	fmt.Println(result)
}

func intersect(nums1 []int, nums2 []int) []int {
	result := []int{}
	numsMap := make(map[int]int, len(nums1))

	for _, value := range nums1 {
		numsMap[value]++
	}

	for _, value := range nums2 {
		fmt.Println(value)

		if numsMap[value] > 0 {
			result = append(result, value)
			numsMap[value]--
		}
	}

	return result
}
