package main

import "fmt"

func main() {
	testCase := []int{-2, 0, 3, -5, 2, -1}

	nums := Constructor(testCase)

	result := nums.SumRange(2, 5)

	fmt.Println(result)
}

type NumArray struct {
	sum []int
}

func Constructor(nums []int) NumArray {
	res := NumArray{make([]int, len(nums))}
	if len(nums) == 0 {
		return res
	}

	res.sum[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		res.sum[i] = res.sum[i-1] + nums[i]
	}

	return res
}

func (this *NumArray) SumRange(i int, j int) int {
	if i == 0 {
		return this.sum[j]
	} else {
		return this.sum[j] - this.sum[i-1]
	}
}
