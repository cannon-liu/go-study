package main

import "fmt"

func countSort(nums []int) []int {
	numLen := len(nums)
	if numLen <= 1 {
		return nums
	}
	maxValue := nums[0]
	for i := 1; i < numLen; i++ {
		if maxValue < nums[i] {
			maxValue = nums[i]
		}
	}
	bucket := make([]int, maxValue+1)
	for i := 0; i < maxValue+1; i++ {
		bucket[i] = 0
	}
	for i := 0; i < numLen; i++ {
		index := nums[i]
		bucket[index]++
	}
	fmt.Println(bucket, " ---- ", len(bucket))
	index := 0
	for i := 0; i < len(bucket); i++ {
		for bucket[i] > 0 {
			fmt.Println(index, i)
			nums[index] = i
			bucket[i]--
			index++
		}
	}
	return nums
}

func main() {
	data := []int{7, 3, 6, 6, 8, 5, 2}
	fmt.Println(data)
	result := countSort(data)
	fmt.Println(result)
}
