package main

import "fmt"

func leastInterval(tasks []byte, n int) int {
	dict := make(map[byte]int)
	for _, task := range tasks {
		_, ok := dict[task]
		if ok {
			dict[task]++
		} else {
			dict[task] = 1
		}
	}

	max := 0
	for _, value := range dict {
		if value > max {
			max = value
		}
	}
	maxCount := 0
	for _, v := range dict {
		if v == max {
			maxCount++
		}
	}

	if len(tasks) > (max-1)*(n+1)+maxCount {
		return len(tasks)
	}
	return (max-1)*(n+1) + maxCount
}

func main() {
	fmt.Println(leastInterval([]byte{'A', 'A', 'A', 'A', 'A', 'A', 'B', 'C', 'D', 'E', 'F', 'G'}, 2))
}
