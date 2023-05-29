package main

import (
	"fmt"
)

type Max struct {
	Max     int64
	NoValue bool
}

func maxStrength(nums []int) int64 {
	var inner func(Max, int) Max
	inner = func(max Max, x int) Max {
		if x == len(nums) {
			return max
		}

		currMax := max.Max

		for i := x; i < len(nums); i++ {
			var curr int64 = 0
			if max.NoValue {
				curr = inner(Max{Max: int64(nums[i])}, i+1).Max
			} else {
				curr = inner(Max{Max: max.Max * int64(nums[i])}, i+1).Max
			}
			if curr > currMax {
				currMax = curr
			}
		}
		return Max{Max: currMax}
	}
	return inner(Max{NoValue: true, Max: -9223372036854775808}, 0).Max
}

func main() {
	fmt.Println(maxStrength([]int{-9}))
}
