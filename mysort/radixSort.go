package main

import (
	"fmt"
)

func radixSort(nums []int) []int {
	numLen := len(nums)
	if numLen <= 1 {
		return nums
	}
	mod := 10
	dev := 1
	maxValue := nums[0]
	maxDigit := 0
	for i := 1; i < numLen; i++ {
		if maxValue < nums[i] {
			maxValue = nums[i]
		}
	}
	fmt.Println("maxValue is ", maxValue)
	for i := 1; i < maxValue; i *= 10 {
		maxDigit++
	}
	fmt.Println("maxDigit is ", maxDigit)
	counter := make([][]int, 10)
	fmt.Println(counter)
	for i := 0; i < maxDigit; i++ {
		for j := 0; j < numLen; j++ {
			bucket := nums[j] % mod / dev
			counter[bucket] = append(counter[bucket], nums[j])
		}
		fmt.Println(counter)
		pos := 0
		for j := 0; j < len(counter); j++ {
			for k := 0; k < len(counter[j]); k++ {
				nums[pos] = counter[j][k]
				pos++
			}
		}
		dev *= 10
		mod *= 10
		for i := 0; i < 10; i++ {
			counter[i] = make([]int, 0)
		}
	}
	return nums
}

func main() {
	data := []int{36, 50, 2, 47, 4, 46, 5, 39, 25, 27, 38, 3, 8, 44, 19, 26}
	fmt.Println(data)
	result := radixSort(data)
	fmt.Println(result)
}
