package main

import (
	"fmt"
	"math"
)

func shellSort(nums []int) []int {
	numLen := len(nums)
	if numLen <= 1 {
		return nums
	}
	for gap := math.Floor(float64(numLen / 2)); gap > 0; gap = math.Floor(gap / 2) {

		for i := int(gap); i < numLen; i++ {
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
			fmt.Println(nums)
		}
		fmt.Println("-----", gap)

	}

	return nums
}

func main() {
	data := []int{7, 3, 6, 6, 8, 5, 2}
	fmt.Println(data)
	result := shellSort(data)
	fmt.Println(result)
}
