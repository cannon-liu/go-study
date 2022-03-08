package main

import (
	"fmt"
	"math"
)

func bucketSort(nums []int, bucketSize int) []int {
	numLen := len(nums)
	if numLen <= 1 {
		return nums
	}
	minValue := nums[0]
	maxValue := nums[0]
	for i := 1; i < numLen; i++ {
		if maxValue < nums[i] {
			maxValue = nums[i]
		}
	}
	for i := 1; i < numLen; i++ {
		if minValue > nums[i] {
			minValue = nums[i]
		}
	}
	DefaultBucketSize := 5
	if bucketSize <= 0 {
		bucketSize = DefaultBucketSize
	}
	bucketCount := int(math.Floor(float64((maxValue-minValue)*1.0/bucketSize)) + 1)
	fmt.Println("bucket count ", bucketCount)
	buckets := make([][]int, bucketCount)
	for i := 0; i < numLen; i++ {
		index := int(math.Floor(float64((nums[i] - minValue) * 1.0 / bucketSize)))
		buckets[index] = append(buckets[index], nums[i])
	}
	data := make([]int, 0)
	for i := 0; i < bucketCount; i++ {
		fmt.Println(buckets[i])
		insertionSort2(buckets[i])
		data = append(data, buckets[i]...)
	}
	return data
}

func insertionSort2(nums []int) []int {
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
	data := []int{78, 17, 39, 26, 72, 94, 21, 12, 23, 68}
	fmt.Println(data)
	result := bucketSort(data, 5)
	fmt.Println(result)
}
