package main

import (
	"fmt"
)

func quickSort(nums []int) []int {
	numLen := len(nums)
	if numLen <= 1 {
		return nums
	}
	left := 0
	right := numLen - 1
	_quickSort(nums, left, right)
	return nums
}

func _quickSort(nums []int, left, right int) {
	var partitionIndex int
	if left < right {
		partitionIndex = partition(nums, left, right)
		_quickSort(nums, left, partitionIndex-1)
		_quickSort(nums, partitionIndex+1, right)
	}
}

func partition(nums []int, left, right int) int {
	pivot := left
	index := pivot + 1
	for i := index; i <= right; i++ {
		if nums[i] < nums[pivot] {
			swap(nums, i, index)
			index++
		}
	}
	swap(nums, pivot, index-1)
	return index - 1
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
	fmt.Println(nums)
}

func main() {
	data := []int{7, 3, 6, 6, 8, 5, 2}
	result := quickSort(data)
	fmt.Println(data)
	fmt.Println(result)
}
