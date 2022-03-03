package main

import "fmt"

func insertionSort(nums []int) []int {
	numLen := len(nums)
	if numLen <= 1 {
		return nums
	}
	for i := 1; i < numLen; i++ {
		preIndex := i - 1
		current := nums[i]
		for {
			if preIndex >= 0 && nums[preIndex] > current {
				nums[preIndex+1] = nums[preIndex]
				preIndex--
			} else {
				break
			}
		}
		nums[preIndex+1] = current
	}
	return nums
}

func main() {
	data := []int{7, 3, 6, 6, 8, 5, 2}
	fmt.Println(data)
	result := insertionSort(data)
	fmt.Println(result)
}
