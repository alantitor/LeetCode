# 136 Single Number

## Tow Layer For Loop

Trace the input array and compare each elements to find the number which is not duplicate.

```go
func singleNumber(nums []int) int {
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
```

## Map

If element is not existing in map, then add it into map.

If element is existing in map, then delete it from map.

At the end, the map will remain one element which is not duplicate in the input array.

```go
func singleNumber(nums []int) int {
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
```

## XOR Operation

Let we analyse the input array. The data have two models.

1. elements exists two times.
2. element is not duplicate.

```go
func singleNumber(nums []int) int {
	var ret int

	for i := range nums {
		ret ^= nums[i]
	}

	return ret
}
```

