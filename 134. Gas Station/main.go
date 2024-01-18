package main

import "fmt"

func canCompleteCircuit(gas []int, cost []int) int {
	arr := make([]int, len(gas))
	for i, _ := range gas {
		arr[i] = gas[i] - cost[i]
	}
	fmt.Println(arr)
	start := 0
	sum := 0
	for i := 0; start <= len(gas); i++ {
		sum += gas[i%len(gas)] - cost[i%len(cost)]
		for sum < 0 {
			sum -= gas[start%len(gas)] - cost[start%len(cost)]
			start++
		}
		if i == start+len(gas) {
			return start
		}
	}
	return -1
}

func main() {
	fmt.Println(canCompleteCircuit([]int{1, 5, 3, 4, 5}, []int{3, 4, 5, 4, 3}))
}
