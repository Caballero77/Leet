package main

import (
	"fmt"
	"sort"
)

func carPooling(trips [][]int, capacity int) bool {
	dict := make(map[int]int)

	for i := 0; i < len(trips); i++ {
		trip := trips[i]
		v, _ := dict[trip[1]]
		dict[trip[1]] = v + trip[0]

		v, _ = dict[trip[2]]
		dict[trip[2]] = v - trip[0]
	}

	keys := make([]int, len(dict))
	for key, _ := range dict {
		keys = append(keys, key)
	}

	sort.Ints(keys)

	for _, key := range keys {
		capacity -= dict[key]

		if capacity < 0 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(carPooling([][]int{{9, 3, 4}, {9, 1, 7}, {4, 2, 4}, {7, 4, 5}}, 23))
}
