package main

import (
	"fmt"
	"github.com/golang/go/src/pkg/math"
)

func main() {
	testNums := [...]int{11, 2, 3, 1, 2, 3,}
	testK := 3

	result := containsNearbyDuplicate(testNums[:], testK)

	fmt.Println(result)
}

func containsNearbyDuplicate(nums []int, k int) bool {
	result := false
	cache := make(map[int]int, len(nums))

	for index, value := range nums {
		if cv, ok := cache[value]; ok {
			if int(math.Abs(float64(index-cv))) <= k {
				result = true
				break
			} else {
				cache[value] = index
			}
		} else {
			cache[value] = index
		}

	}

	return result
}
