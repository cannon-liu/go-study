package main

import "fmt"

func selectionSort(nums []int) []int {
	numLen := len(nums)
	if numLen <= 1 {
		return nums
	}
	for i := numLen; i > 1; i-- {
		maxIndex := 0
		for j := 0; j < i; j++ {
			if nums[maxIndex] < nums[j] {
				maxIndex = j
			}
		}
		nums[i-1], nums[maxIndex] = nums[maxIndex], nums[i-1]
	}
	return nums
}

func main() {
	data := []int{7, 3, 6, 6, 8, 5, 2}
	result := selectionSort(data)
	fmt.Println(data)
	fmt.Println(result)
}
