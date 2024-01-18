package main

import "fmt"

func findPeakElement(nums []int) int {
	i := len(nums) / 2
	fmt.Println(nums)
	fmt.Println(i)
	if (i+1 >= len(nums) || nums[i] > nums[i+1]) && (i-1 < 0 || nums[i] > nums[i-1]) {
		return i
	}

	if i+1 >= len(nums) || nums[i] > nums[i+1] {
		return findPeakElement(nums[:i])
	} else {
		return i + findPeakElement(nums[i:])
	}
}

func main() {
	arr := []int{1}
	fmt.Println(findPeakElement(arr))
}
