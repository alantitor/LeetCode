package main

import "fmt"

func main() {
	data := [...]int{2, 2, 5}

	ans := singleNumber3(data[:])

	fmt.Println(ans)
}

func singleNumber1(nums []int) int {
	ret := 0

	for i := 0; i < len(nums); i++ {
		flag := false

		for j := 0; j < len(nums); j++ {
			if i == j {
				continue
			}

			if nums[i] == nums[j] {
				flag = true
				break
			}
		}

		if !flag {
			ret = nums[i]
			break
		}
	}

	return ret
}

func singleNumber2(nums []int) int {
	var ret int
	m := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		_, isExist := m[nums[i]]

		if isExist {
			delete(m, nums[i])
		} else {
			m[nums[i]] = 1
		}
	}

	for k := range m {
		ret = k
	}

	return ret
}

func singleNumber3(nums []int) int {
	var ret int

	for i := range nums {
		ret ^= nums[i]
	}

	return ret
}
