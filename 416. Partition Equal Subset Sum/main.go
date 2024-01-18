package main

import (
	"fmt"
)

func canPartition(nums []int) bool {
	sum := 0
	max := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if max < nums[i] {
			max = nums[i]
		}
	}
	if sum%2 != 0 || max > sum/2 {
		return false
	}

	dp := make([]bool, sum/2+1)
	dp[0] = true
	for i := 0; i < len(nums); i++ {
		for j := sum / 2; j >= nums[i]; j-- {
			dp[j] = dp[j] || dp[j-nums[i]]
		}
		if dp[sum/2] {
			return true
		}
	}
	return dp[sum/2]
}
func main() {
	fmt.Println(canPartition([]int{1, 5, 5, 11}))
}
