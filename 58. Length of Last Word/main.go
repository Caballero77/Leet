package main

import (
	"fmt"
)

func lengthOfLastWord(s string) int {
	curr := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			curr++
		} else if curr == 0 {
			continue
		} else {
			break
		}
	}
	return curr
}

func main() {
	fmt.Println(lengthOfLastWord(" "))
}
