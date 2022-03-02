package main

import "fmt"

func bubbleSort(nums []int) []int {
	numLen := len(nums)
	if numLen <= 1 {
		return nums
	}
	for i := 0; i < numLen; i++ {
		for j := 0; j < numLen-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}

func main() {
	data := []int{7, 3, 6, 6, 8, 5, 2}
	result := bubbleSort(data)
	fmt.Println(data)
	fmt.Println(result)
}
