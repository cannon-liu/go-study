package main

import (
	"fmt"
	"math"
)

func mergeSort(nums []int) []int {
	numLen := len(nums)
	if numLen <= 1 {
		return nums
	}
	halfLen := int(math.Ceil(float64(numLen / 2)))
	//left := make([]int,int(halfLen))
	//right := make([]int,int(halfLen))
	left := nums[0:halfLen]
	right := nums[halfLen:numLen]
	return merge(mergeSort(left), mergeSort(right))
}

func merge(left, right []int) []int {
	data := make([]int, 0)
	leftLength := len(left)
	rightLength := len(right)
	leftIndex := 0
	rightIndex := 0
	for {
		if leftIndex >= leftLength || rightIndex >= rightLength {
			break
		}
		if left[leftIndex] < right[rightIndex] {
			data = append(data, left[leftIndex])
			leftIndex++
		} else {
			data = append(data, right[rightIndex])
			rightIndex++
		}
	}
	for {
		if leftIndex >= leftLength {
			break
		}
		data = append(data, left[leftIndex])
		leftIndex++
	}
	for {
		if rightIndex >= rightLength {
			break
		}
		data = append(data, right[rightIndex])
		rightIndex++
	}
	return data
}

func main() {
	data := []int{7, 3, 6, 6, 8, 5, 2}
	fmt.Println(data)
	result := mergeSort(data)
	fmt.Println(result)
}
